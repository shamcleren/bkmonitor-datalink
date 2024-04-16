// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	elastic "github.com/olivere/elastic/v7"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/prompb"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/storage"
	"github.com/prometheus/prometheus/storage/remote"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/consul"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/influxdb/decoder"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/log"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/metadata"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/pool"
)

type Instance struct {
	ctx    context.Context
	wg     sync.WaitGroup
	client *elastic.Client

	lock sync.Mutex

	timeout time.Duration
	maxSize int
}

func (i *Instance) QueryRange(ctx context.Context, promql string, start, end time.Time, step time.Duration) (promql.Matrix, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Instance) Query(ctx context.Context, qs string, end time.Time) (promql.Vector, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Instance) QueryExemplar(ctx context.Context, fields []string, query *metadata.Query, start, end time.Time, matchers ...*labels.Matcher) (*decoder.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Instance) LabelNames(ctx context.Context, query *metadata.Query, start, end time.Time, matchers ...*labels.Matcher) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Instance) LabelValues(ctx context.Context, query *metadata.Query, name string, start, end time.Time, matchers ...*labels.Matcher) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (i *Instance) Series(ctx context.Context, query *metadata.Query, start, end time.Time, matchers ...*labels.Matcher) storage.SeriesSet {
	//TODO implement me
	panic("implement me")
}

func (i *Instance) GetInstanceType() string {
	return consul.ElasticsearchStorageType
}

type InstanceOption struct {
	Url        string
	Username   string
	Password   string
	MaxSize    int
	MaxRouting int
	Timeout    time.Duration
}

type queryOption struct {
	index string
	start int64
	end   int64

	query *metadata.Query
}

type indexOpt struct {
	tableID string
	start   int64
	end     int64
}

var TimeSeriesResultPool sync.Pool

func init() {
	TimeSeriesResultPool.New = func() any {
		return &TimeSeriesResult{}
	}
}

func NewInstance(ctx context.Context, opt *InstanceOption) (*Instance, error) {
	if opt.Url == "" || opt.Username == "" || opt.Password == "" {
		return nil, errors.New("empty es client options")
	}
	ins := &Instance{
		ctx:     ctx,
		timeout: opt.Timeout,
		maxSize: opt.MaxSize,
	}

	cli, err := elastic.NewClient(
		elastic.SetURL(opt.Url),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(opt.Username, opt.Password),
	)
	if err != nil {
		return nil, err
	}

	if opt.MaxRouting > 0 {
		err = pool.Tune(opt.MaxRouting)
		if err != nil {
			return nil, err
		}
	}

	ins.client = cli
	return ins, nil
}

func (i *Instance) getAlias(opt *indexOpt) (indexes []string) {
	for ti := opt.start; ti <= opt.end; ti += int64((time.Hour * 24).Seconds()) {
		index := fmt.Sprintf("%s_%s_read", opt.tableID, time.Unix(ti, 0).Format("20060102"))
		indexes = append(indexes, index)
	}

	return
}

func (i *Instance) esAggQuery(ctx context.Context, qo *queryOption, rets chan<- *TimeSeriesResult) error {
	var (
		ret = TimeSeriesResultPool.Get().(*TimeSeriesResult)
		qb  = qo.query
	)

	// 只做聚合计算
	if len(qb.AggregateMethodList) == 0 {
		return nil
	}

	defer func() {
		rets <- ret
	}()

	fact := NewFactory(qb.DataSource)
	query, err := fact.Query(qo.query)
	if err != nil {
		return err
	}

	queryRange := elastic.NewRangeQuery(Timestamp).Gte(qo.start).Lt(qo.end).Format(TimeFormat)
	dslQuery := elastic.NewBoolQuery().Filter(queryRange, query)

	ss := i.client.Search().
		Size(0).
		From(0).
		Index(qo.index).
		Query(dslQuery).
		Sort(Timestamp, true)

	if len(qb.Source) > 0 {
		fetchSource := elastic.NewFetchSourceContext(true)
		fetchSource.Include(qb.Source...)
		ss = ss.FetchSourceContext(fetchSource)
	}

	aggs, err := fact.Aggs(qo.query)
	if err != nil {
		return err
	}

	ss.Aggregation(aggs.Agg().Name, aggs.Agg().Agg)

	fetchSource := elastic.NewFetchSourceContext(true)
	fetchSource.Include("")
	ss = ss.FetchSourceContext(fetchSource)

	sr, err := ss.Do(ctx)
	if err != nil {
		return nil
	}

	res, err := dataFormat(aggs.Aggs, sr.Aggregations, fact.Relabel)
	if err != nil {
		return err
	}

	log.Debugf(ctx, "es agg query %d, %d, %s, result: %s", qo.start, qo.end, qo.index, res.String())

	ret = &TimeSeriesResult{
		TimeSeriesMap: res.TimeSeriesMap,
	}
	return nil
}

func (i *Instance) Close() {
	i.wg.Wait()
	i.ctx = nil
	i.client = nil
}

func (i *Instance) getIndexes(ctx context.Context, aliases []string) ([]string, error) {
	catAlias, err := i.client.CatAliases().Alias(aliases...).Do(ctx)
	if err != nil {
		return nil, err
	}

	indexMap := make(map[string]struct{}, 0)
	for _, a := range catAlias {
		indexMap[a.Index] = struct{}{}
	}
	indexes := make([]string, 0, len(indexMap))
	for idx := range indexMap {
		indexes = append(indexes, idx)
	}

	return indexes, nil
}

func (i *Instance) indexOption(ctx context.Context, index string) (docCount int64, storeSize int64, err error) {
	cats, err := i.client.CatIndices().Index(index).Do(ctx)
	if err != nil {
		return
	}
	for _, c := range cats {
		docCount = int64(c.DocsCount)
		storeSize, err = parseSizeString(c.StoreSize)
		if err != nil {
			return
		}
		break
	}

	return
}

func (i *Instance) makeQueryOption(ctx context.Context, query *metadata.Query, start, end int64) (indexQueryOpts []*queryOption, err error) {
	aliases := make([]string, 0)
	for ti := start; ti <= end; ti += (time.Hour * 24).Milliseconds() {
		alias := fmt.Sprintf("%s_%s_read", query.DB, time.UnixMilli(ti).Format("20060102"))
		aliases = append(aliases, alias)
	}
	indexes, err := i.getIndexes(ctx, aliases)
	if err != nil {
		return
	}
	if len(indexes) == 0 {
		err = fmt.Errorf("empty index with tableID %+v", query.TableID)
		return
	}

	indexQueryOpts = make([]*queryOption, 0)

	for _, index := range indexes {
		docCount, storeSize, err1 := i.indexOption(ctx, index)
		if err1 != nil {
			err = err1
			return
		}
		qs, err1 := newRangeSegment(ctx, &querySegmentOption{
			start:     start,
			end:       end,
			interval:  query.TimeAggregation.WindowDuration.Milliseconds(),
			docCount:  docCount,
			storeSize: storeSize,
		})
		if err1 != nil {
			err = err1
			return
		}

		for _, l := range qs.list {
			nqo := &queryOption{
				query: query,
			}
			nqo.index = index
			nqo.start = l[0]
			nqo.end = l[1]
			indexQueryOpts = append(indexQueryOpts, nqo)
		}
		qs.close()
	}

	return
}

func (i *Instance) mergeTimeSeries(rets chan *TimeSeriesResult) (storage.SeriesSet, error) {
	seriesMap := make(map[string]*prompb.TimeSeries)

	for ret := range rets {
		if len(ret.TimeSeriesMap) == 0 {
			continue
		}

		for key, ts := range ret.TimeSeriesMap {
			if _, ok := seriesMap[key]; !ok {
				seriesMap[key] = &prompb.TimeSeries{
					Labels:  ts.GetLabels(),
					Samples: make([]prompb.Sample, 0),
				}
			}

			seriesMap[key].Samples = append(seriesMap[key].Samples, ts.Samples...)
		}

		ret.TimeSeriesMap = nil
		TimeSeriesResultPool.Put(ret)
	}

	qr := &prompb.QueryResult{
		Timeseries: make([]*prompb.TimeSeries, 0, len(seriesMap)),
	}
	for _, ts := range seriesMap {
		sort.Slice(ts.Samples, func(i, j int) bool {
			return ts.Samples[i].GetTimestamp() < ts.Samples[j].GetTimestamp()
		})

		qr.Timeseries = append(qr.Timeseries, ts)
	}

	return remote.FromQueryResult(false, qr), nil
}

// QueryRaw 查询原始数据
func (i *Instance) QueryRaw(
	ctx context.Context,
	query *metadata.Query,
	hints *storage.SelectHints,
	matchers ...*labels.Matcher,
) storage.SeriesSet {
	start := hints.Start
	end := hints.End

	// 获取 window 对齐开始时间
	if query.TimeAggregation == nil {
		err := fmt.Errorf("empty time aggregation with %+v", query)
		return storage.ErrSeriesSet(err)
	}
	window := query.TimeAggregation.WindowDuration
	if window.Milliseconds() > 0 {
		start = intMathFloor(start, window.Milliseconds()) * window.Milliseconds()
	}

	indexQueryOpts, err := i.makeQueryOption(ctx, query, start, end)
	if err != nil {
		return storage.ErrSeriesSet(err)
	}

	rets := make(chan *TimeSeriesResult, len(indexQueryOpts))

	go func() {
		var wg sync.WaitGroup
		defer func() {
			wg.Wait()
			close(rets)
		}()
		for _, qo := range indexQueryOpts {
			wg.Add(1)
			qo := qo
			err = pool.Submit(func() {
				defer func() {
					wg.Done()
				}()
				err = i.esAggQuery(ctx, qo, rets)
			})
		}
	}()

	if err != nil {
		return storage.ErrSeriesSet(err)
	}

	set, err := i.mergeTimeSeries(rets)
	if err != nil {
		return storage.ErrSeriesSet(err)
	}

	return set
}
