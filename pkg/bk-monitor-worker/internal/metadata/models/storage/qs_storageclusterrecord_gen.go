// Code generated by go-queryset. DO NOT EDIT.
package storage

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set ClusterRecordQuerySet

// ClusterRecordQuerySet is an queryset type for ClusterRecord
type ClusterRecordQuerySet struct {
	db *gorm.DB
}

// NewClusterRecordQuerySet constructs new ClusterRecordQuerySet
func NewClusterRecordQuerySet(db *gorm.DB) ClusterRecordQuerySet {
	return ClusterRecordQuerySet{
		db: db.Model(&ClusterRecord{}),
	}
}

func (qs ClusterRecordQuerySet) w(db *gorm.DB) ClusterRecordQuerySet {
	return NewClusterRecordQuerySet(db)
}

func (qs ClusterRecordQuerySet) Select(fields ...ClusterRecordDBSchemaField) ClusterRecordQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *ClusterRecord) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *ClusterRecord) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) All(ret *[]ClusterRecord) error {
	return qs.db.Find(ret).Error
}

// ClusterIDEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) ClusterIDEq(clusterID int64) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("cluster_id = ?", clusterID))
}

// ClusterIDGt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) ClusterIDGt(clusterID int64) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("cluster_id > ?", clusterID))
}

// ClusterIDGte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) ClusterIDGte(clusterID int64) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("cluster_id >= ?", clusterID))
}

// ClusterIDIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) ClusterIDIn(clusterID ...int64) ClusterRecordQuerySet {
	if len(clusterID) == 0 {
		qs.db.AddError(errors.New("must at least pass one clusterID in ClusterIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("cluster_id IN (?)", clusterID))
}

// ClusterIDLt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) ClusterIDLt(clusterID int64) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("cluster_id < ?", clusterID))
}

// ClusterIDLte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) ClusterIDLte(clusterID int64) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("cluster_id <= ?", clusterID))
}

// ClusterIDNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) ClusterIDNe(clusterID int64) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("cluster_id != ?", clusterID))
}

// ClusterIDNotIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) ClusterIDNotIn(clusterID ...int64) ClusterRecordQuerySet {
	if len(clusterID) == 0 {
		qs.db.AddError(errors.New("must at least pass one clusterID in ClusterIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("cluster_id NOT IN (?)", clusterID))
}

// Count is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreateTimeEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreateTimeEq(createTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("create_time = ?", createTime))
}

// CreateTimeGt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreateTimeGt(createTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("create_time > ?", createTime))
}

// CreateTimeGte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreateTimeGte(createTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("create_time >= ?", createTime))
}

// CreateTimeLt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreateTimeLt(createTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("create_time < ?", createTime))
}

// CreateTimeLte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreateTimeLte(createTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("create_time <= ?", createTime))
}

// CreateTimeNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreateTimeNe(createTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("create_time != ?", createTime))
}

// CreatorEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorEq(creator string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("creator = ?", creator))
}

// CreatorGt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorGt(creator string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("creator > ?", creator))
}

// CreatorGte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorGte(creator string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("creator >= ?", creator))
}

// CreatorIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorIn(creator ...string) ClusterRecordQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator IN (?)", creator))
}

// CreatorLike is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorLike(creator string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("creator LIKE ?", creator))
}

// CreatorLt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorLt(creator string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("creator < ?", creator))
}

// CreatorLte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorLte(creator string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("creator <= ?", creator))
}

// CreatorNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorNe(creator string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("creator != ?", creator))
}

// CreatorNotIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorNotIn(creator ...string) ClusterRecordQuerySet {
	if len(creator) == 0 {
		qs.db.AddError(errors.New("must at least pass one creator in CreatorNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("creator NOT IN (?)", creator))
}

// CreatorNotlike is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) CreatorNotlike(creator string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("creator NOT LIKE ?", creator))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) Delete() error {
	return qs.db.Delete(ClusterRecord{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(ClusterRecord{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(ClusterRecord{})
	return db.RowsAffected, db.Error
}

// DeleteTimeEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteTimeEq(deleteTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("delete_time = ?", deleteTime))
}

// DeleteTimeGt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteTimeGt(deleteTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("delete_time > ?", deleteTime))
}

// DeleteTimeGte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteTimeGte(deleteTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("delete_time >= ?", deleteTime))
}

// DeleteTimeIsNotNull is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteTimeIsNotNull() ClusterRecordQuerySet {
	return qs.w(qs.db.Where("delete_time IS NOT NULL"))
}

// DeleteTimeIsNull is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteTimeIsNull() ClusterRecordQuerySet {
	return qs.w(qs.db.Where("delete_time IS NULL"))
}

// DeleteTimeLt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteTimeLt(deleteTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("delete_time < ?", deleteTime))
}

// DeleteTimeLte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteTimeLte(deleteTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("delete_time <= ?", deleteTime))
}

// DeleteTimeNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DeleteTimeNe(deleteTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("delete_time != ?", deleteTime))
}

// DisableTimeEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DisableTimeEq(disableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("disable_time = ?", disableTime))
}

// DisableTimeGt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DisableTimeGt(disableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("disable_time > ?", disableTime))
}

// DisableTimeGte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DisableTimeGte(disableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("disable_time >= ?", disableTime))
}

// DisableTimeIsNotNull is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DisableTimeIsNotNull() ClusterRecordQuerySet {
	return qs.w(qs.db.Where("disable_time IS NOT NULL"))
}

// DisableTimeIsNull is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DisableTimeIsNull() ClusterRecordQuerySet {
	return qs.w(qs.db.Where("disable_time IS NULL"))
}

// DisableTimeLt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DisableTimeLt(disableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("disable_time < ?", disableTime))
}

// DisableTimeLte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DisableTimeLte(disableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("disable_time <= ?", disableTime))
}

// DisableTimeNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) DisableTimeNe(disableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("disable_time != ?", disableTime))
}

// EnableTimeEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) EnableTimeEq(enableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("enable_time = ?", enableTime))
}

// EnableTimeGt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) EnableTimeGt(enableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("enable_time > ?", enableTime))
}

// EnableTimeGte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) EnableTimeGte(enableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("enable_time >= ?", enableTime))
}

// EnableTimeIsNotNull is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) EnableTimeIsNotNull() ClusterRecordQuerySet {
	return qs.w(qs.db.Where("enable_time IS NOT NULL"))
}

// EnableTimeIsNull is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) EnableTimeIsNull() ClusterRecordQuerySet {
	return qs.w(qs.db.Where("enable_time IS NULL"))
}

// EnableTimeLt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) EnableTimeLt(enableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("enable_time < ?", enableTime))
}

// EnableTimeLte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) EnableTimeLte(enableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("enable_time <= ?", enableTime))
}

// EnableTimeNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) EnableTimeNe(enableTime time.Time) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("enable_time != ?", enableTime))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) GetUpdater() ClusterRecordUpdater {
	return NewClusterRecordUpdater(qs.db)
}

// IsCurrentEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) IsCurrentEq(isCurrent bool) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("is_current = ?", isCurrent))
}

// IsCurrentIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) IsCurrentIn(isCurrent ...bool) ClusterRecordQuerySet {
	if len(isCurrent) == 0 {
		qs.db.AddError(errors.New("must at least pass one isCurrent in IsCurrentIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("is_current IN (?)", isCurrent))
}

// IsCurrentNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) IsCurrentNe(isCurrent bool) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("is_current != ?", isCurrent))
}

// IsCurrentNotIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) IsCurrentNotIn(isCurrent ...bool) ClusterRecordQuerySet {
	if len(isCurrent) == 0 {
		qs.db.AddError(errors.New("must at least pass one isCurrent in IsCurrentNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("is_current NOT IN (?)", isCurrent))
}

// IsDeletedEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) IsDeletedEq(isDeleted bool) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("is_deleted = ?", isDeleted))
}

// IsDeletedIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) IsDeletedIn(isDeleted ...bool) ClusterRecordQuerySet {
	if len(isDeleted) == 0 {
		qs.db.AddError(errors.New("must at least pass one isDeleted in IsDeletedIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("is_deleted IN (?)", isDeleted))
}

// IsDeletedNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) IsDeletedNe(isDeleted bool) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("is_deleted != ?", isDeleted))
}

// IsDeletedNotIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) IsDeletedNotIn(isDeleted ...bool) ClusterRecordQuerySet {
	if len(isDeleted) == 0 {
		qs.db.AddError(errors.New("must at least pass one isDeleted in IsDeletedNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("is_deleted NOT IN (?)", isDeleted))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) Limit(limit int) ClusterRecordQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) Offset(offset int) ClusterRecordQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ClusterRecordQuerySet) One(ret *ClusterRecord) error {
	return qs.db.First(ret).Error
}

// OrderAscByClusterID is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByClusterID() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("cluster_id ASC"))
}

// OrderAscByCreateTime is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByCreateTime() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("create_time ASC"))
}

// OrderAscByCreator is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByCreator() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("creator ASC"))
}

// OrderAscByDeleteTime is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByDeleteTime() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("delete_time ASC"))
}

// OrderAscByDisableTime is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByDisableTime() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("disable_time ASC"))
}

// OrderAscByEnableTime is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByEnableTime() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("enable_time ASC"))
}

// OrderAscByIsCurrent is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByIsCurrent() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("is_current ASC"))
}

// OrderAscByIsDeleted is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByIsDeleted() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("is_deleted ASC"))
}

// OrderAscByTableID is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderAscByTableID() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("table_id ASC"))
}

// OrderDescByClusterID is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByClusterID() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("cluster_id DESC"))
}

// OrderDescByCreateTime is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByCreateTime() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("create_time DESC"))
}

// OrderDescByCreator is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByCreator() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("creator DESC"))
}

// OrderDescByDeleteTime is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByDeleteTime() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("delete_time DESC"))
}

// OrderDescByDisableTime is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByDisableTime() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("disable_time DESC"))
}

// OrderDescByEnableTime is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByEnableTime() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("enable_time DESC"))
}

// OrderDescByIsCurrent is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByIsCurrent() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("is_current DESC"))
}

// OrderDescByIsDeleted is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByIsDeleted() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("is_deleted DESC"))
}

// OrderDescByTableID is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) OrderDescByTableID() ClusterRecordQuerySet {
	return qs.w(qs.db.Order("table_id DESC"))
}

// TableIDEq is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDEq(tableID string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("table_id = ?", tableID))
}

// TableIDGt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDGt(tableID string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("table_id > ?", tableID))
}

// TableIDGte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDGte(tableID string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("table_id >= ?", tableID))
}

// TableIDIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDIn(tableID ...string) ClusterRecordQuerySet {
	if len(tableID) == 0 {
		qs.db.AddError(errors.New("must at least pass one tableID in TableIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("table_id IN (?)", tableID))
}

// TableIDLike is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDLike(tableID string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("table_id LIKE ?", tableID))
}

// TableIDLt is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDLt(tableID string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("table_id < ?", tableID))
}

// TableIDLte is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDLte(tableID string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("table_id <= ?", tableID))
}

// TableIDNe is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDNe(tableID string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("table_id != ?", tableID))
}

// TableIDNotIn is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDNotIn(tableID ...string) ClusterRecordQuerySet {
	if len(tableID) == 0 {
		qs.db.AddError(errors.New("must at least pass one tableID in TableIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("table_id NOT IN (?)", tableID))
}

// TableIDNotlike is an autogenerated method
// nolint: dupl
func (qs ClusterRecordQuerySet) TableIDNotlike(tableID string) ClusterRecordQuerySet {
	return qs.w(qs.db.Where("table_id NOT LIKE ?", tableID))
}

// SetClusterID is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetClusterID(clusterID int64) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.ClusterID)] = clusterID
	return u
}

// SetCreateTime is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetCreateTime(createTime time.Time) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.CreateTime)] = createTime
	return u
}

// SetCreator is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetCreator(creator string) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.Creator)] = creator
	return u
}

// SetDeleteTime is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetDeleteTime(deleteTime *time.Time) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.DeleteTime)] = deleteTime
	return u
}

// SetDisableTime is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetDisableTime(disableTime *time.Time) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.DisableTime)] = disableTime
	return u
}

// SetEnableTime is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetEnableTime(enableTime *time.Time) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.EnableTime)] = enableTime
	return u
}

// SetIsCurrent is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetIsCurrent(isCurrent bool) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.IsCurrent)] = isCurrent
	return u
}

// SetIsDeleted is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetIsDeleted(isDeleted bool) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.IsDeleted)] = isDeleted
	return u
}

// SetTableID is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) SetTableID(tableID string) ClusterRecordUpdater {
	u.fields[string(ClusterRecordDBSchema.TableID)] = tableID
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ClusterRecordUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set ClusterRecordQuerySet

// ===== BEGIN of ClusterRecord modifiers

// ClusterRecordDBSchemaField describes database schema field. It requires for method 'Update'
type ClusterRecordDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ClusterRecordDBSchemaField) String() string {
	return string(f)
}

// ClusterRecordDBSchema stores db field names of ClusterRecord
var ClusterRecordDBSchema = struct {
	TableID     ClusterRecordDBSchemaField
	ClusterID   ClusterRecordDBSchemaField
	IsDeleted   ClusterRecordDBSchemaField
	IsCurrent   ClusterRecordDBSchemaField
	Creator     ClusterRecordDBSchemaField
	CreateTime  ClusterRecordDBSchemaField
	EnableTime  ClusterRecordDBSchemaField
	DisableTime ClusterRecordDBSchemaField
	DeleteTime  ClusterRecordDBSchemaField
}{

	TableID:     ClusterRecordDBSchemaField("table_id"),
	ClusterID:   ClusterRecordDBSchemaField("cluster_id"),
	IsDeleted:   ClusterRecordDBSchemaField("is_deleted"),
	IsCurrent:   ClusterRecordDBSchemaField("is_current"),
	Creator:     ClusterRecordDBSchemaField("creator"),
	CreateTime:  ClusterRecordDBSchemaField("create_time"),
	EnableTime:  ClusterRecordDBSchemaField("enable_time"),
	DisableTime: ClusterRecordDBSchemaField("disable_time"),
	DeleteTime:  ClusterRecordDBSchemaField("delete_time"),
}

// Update updates ClusterRecord fields by primary key
// nolint: dupl
func (o *ClusterRecord) Update(db *gorm.DB, fields ...ClusterRecordDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"table_id":     o.TableID,
		"cluster_id":   o.ClusterID,
		"is_deleted":   o.IsDeleted,
		"is_current":   o.IsCurrent,
		"creator":      o.Creator,
		"create_time":  o.CreateTime,
		"enable_time":  o.EnableTime,
		"disable_time": o.DisableTime,
		"delete_time":  o.DeleteTime,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update ClusterRecord %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ClusterRecordUpdater is an ClusterRecord updates manager
type ClusterRecordUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewClusterRecordUpdater creates new ClusterRecord updater
// nolint: dupl
func NewClusterRecordUpdater(db *gorm.DB) ClusterRecordUpdater {
	return ClusterRecordUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&ClusterRecord{}),
	}
}

// ===== END of ClusterRecord modifiers

// ===== END of all query sets
