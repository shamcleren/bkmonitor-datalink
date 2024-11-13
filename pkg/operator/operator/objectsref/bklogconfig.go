// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package objectsref

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	loggingv1alpha1 "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/apis/logging/v1alpha1"
	bkversioned "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/client/clientset/versioned"
	bkinformers "github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/client/informers/externalversions"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/common/define"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/common/feature"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/common/k8sutils"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/configs"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/logger"
)

const (
	logConfigTypeStd       = "std_log_config"
	logConfigTypeContainer = "container_log_config"
	logConfigTypeNode      = "node_log_config"
)

const (
	kindBkLogConfig    = "BkLogConfig"
	groupBkLogConfig   = "bk.tencent.com"
	versionBkLogConfig = "v1alpha1"

	resourceBkLogConfigs = "bklogconfigs"
)

type bkLogConfigEntity struct {
	Obj *loggingv1alpha1.BkLogConfig
}

func newBkLogConfigEntity(obj any) *bkLogConfigEntity {
	bkLogConfig, ok := obj.(*loggingv1alpha1.BkLogConfig)
	if !ok {
		logger.Errorf("expected %s type, got %T", kindBkLogConfig, obj)
		return nil
	}

	// check bk env
	env := feature.BkEnv(bkLogConfig.Labels)
	if env != configs.G().BkEnv {
		logger.Warnf("[%s] want bkenv '%s', but got '%s'", kindBkLogConfig, configs.G().BkEnv, env)
		return nil
	}

	return &bkLogConfigEntity{
		Obj: bkLogConfig,
	}
}

func (e *bkLogConfigEntity) UUID() string {
	if e.Obj == nil {
		return ""
	}

	return fmt.Sprintf("%s/%s", e.Obj.Namespace, e.Obj.Name)
}

func (e *bkLogConfigEntity) isVCluster(matcherLabel map[string]string) bool {
	vClusterLabelKey := configs.G().VCluster.LabelKey
	_, ok := matcherLabel[vClusterLabelKey]
	return ok
}

func (e *bkLogConfigEntity) getValues(matcherLabel map[string]string, key string, defaultValue string) string {
	if v, ok := matcherLabel[key]; ok {
		return v
	}
	return defaultValue
}

func (e *bkLogConfigEntity) toLowerEq(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}

func (e *bkLogConfigEntity) getWorkloadName(name string, kind string) string {
	if e.toLowerEq(kind, kindReplicaSet) {
		index := strings.LastIndex(name, "-")
		return name[:index]
	}
	return name
}

func (e *bkLogConfigEntity) MatchWorkloadName(matcherLabels, matcherAnnotations map[string]string, ownerRefs []OwnerRef) bool {
	if e.Obj.Spec.WorkloadName == "" {
		return true
	}

	r, err := regexp.Compile(e.Obj.Spec.WorkloadName)
	if err != nil {
		return false
	}

	var names []string

	if e.isVCluster(matcherLabels) {
		name := e.getValues(matcherAnnotations, configs.G().VCluster.WorkloadNameAnnotationKey, "")
		kind := e.getValues(matcherAnnotations, configs.G().VCluster.WorkloadTypeAnnotationKey, "")
		names = append(names, e.getWorkloadName(name, kind))
	} else {
		for _, ownerReference := range ownerRefs {
			names = append(names, e.getWorkloadName(ownerReference.Name, ownerReference.Kind))
		}
	}

	for _, name := range names {
		if r.MatchString(name) {
			return true
		}
		if e.toLowerEq(name, e.Obj.Spec.WorkloadName) {
			return true
		}
	}
	return false
}

func (e *bkLogConfigEntity) MatchWorkloadType(matcherLabels, matcherAnnotations map[string]string, ownerRefs []OwnerRef) bool {
	var kinds []string

	if e.Obj.Spec.WorkloadType == "" {
		return true
	}

	if e.isVCluster(matcherLabels) {
		kinds = append(kinds, e.getValues(matcherAnnotations, configs.G().VCluster.WorkloadTypeAnnotationKey, ""))
	} else {
		for _, ownerReference := range ownerRefs {
			kinds = append(kinds, ownerReference.Kind)
		}
	}

	for _, kind := range kinds {
		if e.toLowerEq(kind, kindReplicaSet) {
			if e.toLowerEq(e.Obj.Spec.WorkloadType, kindDeployment) {
				return true
			}
		}
		if e.toLowerEq(e.Obj.Spec.WorkloadType, kind) {
			return true
		}
	}
	return false
}

func (e *bkLogConfigEntity) MatchContainerName(containerName string) bool {
	// containerNameMatch empty return true because do not match containerName
	if len(e.Obj.Spec.ContainerNameExclude) != 0 {
		for _, excludeName := range e.Obj.Spec.ContainerNameExclude {
			if excludeName == containerName {
				// containerName is in containerNameExclude, return false
				return false
			}
		}
	}
	if len(e.Obj.Spec.ContainerNameMatch) == 0 {
		return true
	}
	for _, matchContainerName := range e.Obj.Spec.ContainerNameMatch {
		if matchContainerName == containerName {
			return true
		}
	}
	return false
}

func (e *bkLogConfigEntity) MatchAnnotation(matchAnnotations map[string]string) bool {
	selector, err := metav1.LabelSelectorAsSelector(&e.Obj.Spec.AnnotationSelector)
	if err != nil {
		return false
	}

	labelSet := labels.Set(matchAnnotations)
	if !selector.Matches(labelSet) {
		return false
	}

	return true
}

func (e *bkLogConfigEntity) MatchLabel(matchLabels map[string]string) bool {
	selector, err := metav1.LabelSelectorAsSelector(&e.Obj.Spec.LabelSelector)
	if err != nil {
		return false
	}

	labelSet := labels.Set(matchLabels)
	if !selector.Matches(labelSet) {
		return false
	}

	return true
}

// MatchNamespace 判断 namespace 是否匹配上
func (e *bkLogConfigEntity) MatchNamespace(namespace string) bool {
	if e.Obj.Spec.NamespaceSelector.Any {
		return true
	}

	if len(e.Obj.Spec.NamespaceSelector.ExcludeNames) != 0 {
		// 全部不匹配 true，否则为 false
		for _, ns := range e.Obj.Spec.NamespaceSelector.ExcludeNames {
			if ns == namespace {
				return false
			}
		}
		return true
	} else if len(e.Obj.Spec.NamespaceSelector.MatchNames) != 0 {
		// 优先使用 NamespaceSelector 配置，列表中任意一个满足即可
		// 有一个匹配上则为true，否则直接false
		for _, ns := range e.Obj.Spec.NamespaceSelector.MatchNames {
			if ns == namespace {
				return true
			}
		}
		return false
	} else {
		// 其次，使用 Namespace 配置，直接名字匹配
		if e.Obj.Spec.Namespace != "" {
			if e.Obj.Spec.Namespace != namespace {
				return false
			}
			return true
		}
		// 未配置则返回true
		return true
	}
}

type BkLogConfigMap struct {
	lock sync.RWMutex

	entitiesMap map[string]*bkLogConfigEntity
}

func (o *BkLogConfigMap) deleteEntity(e *bkLogConfigEntity) {
	o.lock.Lock()
	defer o.lock.Unlock()
	delete(o.entitiesMap, e.UUID())

	logger.Infof("[%s] delete %s, length: %d", kindBkLogConfig, e.UUID(), len(o.entitiesMap))
}

func (o *BkLogConfigMap) setEntity(e *bkLogConfigEntity) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.entitiesMap[e.UUID()] = e

	logger.Infof("[%s] set %s, length: %d", kindBkLogConfig, e.UUID(), len(o.entitiesMap))
}

func (o *BkLogConfigMap) RangeBkLogConfig(visitFunc func(e *bkLogConfigEntity)) {
	o.lock.RLock()
	defer o.lock.RUnlock()

	for _, e := range o.entitiesMap {
		if visitFunc != nil {
			visitFunc(e)
		}
	}
}

func NewObjectsMap(ctx context.Context, client bkversioned.Interface, resources map[GVRK]struct{}) (*BkLogConfigMap, error) {
	objsMap := &BkLogConfigMap{
		entitiesMap: make(map[string]*bkLogConfigEntity),
	}

	gvrk := GVRK{
		Group:    groupBkLogConfig,
		Version:  versionBkLogConfig,
		Resource: resourceBkLogConfigs,
		Kind:     kindBkLogConfig,
	}

	if _, ok := resources[gvrk]; !ok {
		return objsMap, nil
	}

	factory := bkinformers.NewSharedInformerFactory(client, define.ReSyncPeriod)
	informer := factory.Bk().V1alpha1().BkLogConfigs().Informer()

	informer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				bkLogConfig := newBkLogConfigEntity(obj)
				if bkLogConfig == nil {
					return
				}

				objsMap.setEntity(bkLogConfig)
			},
			UpdateFunc: func(_, newObj interface{}) {
				bkLogConfig := newBkLogConfigEntity(newObj)
				if bkLogConfig == nil {
					return
				}

				objsMap.setEntity(bkLogConfig)
			},
			DeleteFunc: func(obj interface{}) {
				bkLogConfig := newBkLogConfigEntity(obj)
				if bkLogConfig == nil {
					return
				}

				objsMap.deleteEntity(bkLogConfig)
			},
		},
	)
	go informer.Run(ctx.Done())
	logger.Infof("[%s] informer start", kindBkLogConfig)

	synced := k8sutils.WaitForNamedCacheSync(ctx, kindBkLogConfig, informer)
	if !synced {
		return nil, fmt.Errorf("[%s] failed to sync caches", kindBkLogConfig)
	}

	return objsMap, nil
}
