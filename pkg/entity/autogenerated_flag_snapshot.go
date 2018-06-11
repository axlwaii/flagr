// Code generated by go-queryset. DO NOT EDIT.
package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// notest
// ===== BEGIN of all query sets

// ===== BEGIN of query set FlagSnapshotQuerySet

// FlagSnapshotQuerySet is an queryset type for FlagSnapshot
type FlagSnapshotQuerySet struct {
	db *gorm.DB
}

// NewFlagSnapshotQuerySet constructs new FlagSnapshotQuerySet
func NewFlagSnapshotQuerySet(db *gorm.DB) FlagSnapshotQuerySet {
	return FlagSnapshotQuerySet{
		db: db.Model(&FlagSnapshot{}),
	}
}

func (qs FlagSnapshotQuerySet) w(db *gorm.DB) FlagSnapshotQuerySet {
	return NewFlagSnapshotQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) All(ret *[]FlagSnapshot) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *FlagSnapshot) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) CreatedAtEq(createdAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) CreatedAtGt(createdAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) CreatedAtGte(createdAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) CreatedAtLt(createdAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) CreatedAtLte(createdAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) CreatedAtNe(createdAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) Delete() error {
	return qs.db.Delete(FlagSnapshot{}).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *FlagSnapshot) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) DeletedAtEq(deletedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) DeletedAtGt(deletedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) DeletedAtGte(deletedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) DeletedAtIsNotNull() FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) DeletedAtIsNull() FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) DeletedAtLt(deletedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) DeletedAtLte(deletedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) DeletedAtNe(deletedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// FlagIDEq is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIDEq(flagID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("flag_id = ?", flagID))
}

// FlagIDGt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIDGt(flagID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("flag_id > ?", flagID))
}

// FlagIDGte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIDGte(flagID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("flag_id >= ?", flagID))
}

// FlagIDIn is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIDIn(flagID ...uint) FlagSnapshotQuerySet {
	if len(flagID) == 0 {
		qs.db.AddError(errors.New("must at least pass one flagID in FlagIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("flag_id IN (?)", flagID))
}

// FlagIDLt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIDLt(flagID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("flag_id < ?", flagID))
}

// FlagIDLte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIDLte(flagID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("flag_id <= ?", flagID))
}

// FlagIDNe is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIDNe(flagID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("flag_id != ?", flagID))
}

// FlagIDNotIn is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIDNotIn(flagID ...uint) FlagSnapshotQuerySet {
	if len(flagID) == 0 {
		qs.db.AddError(errors.New("must at least pass one flagID in FlagIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("flag_id NOT IN (?)", flagID))
}

// FlagIsNotNull is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIsNotNull() FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("flag IS NOT NULL"))
}

// FlagIsNull is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) FlagIsNull() FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("flag IS NULL"))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) GetUpdater() FlagSnapshotUpdater {
	return NewFlagSnapshotUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) IDEq(ID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) IDGt(ID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) IDGte(ID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) IDIn(ID ...uint) FlagSnapshotQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) IDLt(ID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) IDLte(ID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) IDNe(ID uint) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) IDNotIn(ID ...uint) FlagSnapshotQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) Limit(limit int) FlagSnapshotQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs FlagSnapshotQuerySet) One(ret *FlagSnapshot) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderAscByCreatedAt() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderAscByDeletedAt() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByFlagID is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderAscByFlagID() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("flag_id ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderAscByID() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderAscByUpdatedAt() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderDescByCreatedAt() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderDescByDeletedAt() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByFlagID is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderDescByFlagID() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("flag_id DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderDescByID() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) OrderDescByUpdatedAt() FlagSnapshotQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PreloadFlag is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) PreloadFlag() FlagSnapshotQuerySet {
	return qs.w(qs.db.Preload("Flag"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u FlagSnapshotUpdater) SetCreatedAt(createdAt time.Time) FlagSnapshotUpdater {
	u.fields[string(FlagSnapshotDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u FlagSnapshotUpdater) SetDeletedAt(deletedAt *time.Time) FlagSnapshotUpdater {
	u.fields[string(FlagSnapshotDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetFlagID is an autogenerated method
// nolint: dupl
func (u FlagSnapshotUpdater) SetFlagID(flagID uint) FlagSnapshotUpdater {
	u.fields[string(FlagSnapshotDBSchema.FlagID)] = flagID
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u FlagSnapshotUpdater) SetID(ID uint) FlagSnapshotUpdater {
	u.fields[string(FlagSnapshotDBSchema.ID)] = ID
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u FlagSnapshotUpdater) SetUpdatedAt(updatedAt time.Time) FlagSnapshotUpdater {
	u.fields[string(FlagSnapshotDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUpdatedBy is an autogenerated method
// nolint: dupl
func (u FlagSnapshotUpdater) SetUpdatedBy(updatedBy string) FlagSnapshotUpdater {
	u.fields[string(FlagSnapshotDBSchema.UpdatedBy)] = updatedBy
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u FlagSnapshotUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u FlagSnapshotUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedAtEq(updatedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedAtGt(updatedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedAtGte(updatedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedAtLt(updatedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedAtLte(updatedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedAtNe(updatedAt time.Time) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UpdatedByEq is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedByEq(updatedBy string) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("updated_by = ?", updatedBy))
}

// UpdatedByIn is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedByIn(updatedBy ...string) FlagSnapshotQuerySet {
	if len(updatedBy) == 0 {
		qs.db.AddError(errors.New("must at least pass one updatedBy in UpdatedByIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updated_by IN (?)", updatedBy))
}

// UpdatedByNe is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedByNe(updatedBy string) FlagSnapshotQuerySet {
	return qs.w(qs.db.Where("updated_by != ?", updatedBy))
}

// UpdatedByNotIn is an autogenerated method
// nolint: dupl
func (qs FlagSnapshotQuerySet) UpdatedByNotIn(updatedBy ...string) FlagSnapshotQuerySet {
	if len(updatedBy) == 0 {
		qs.db.AddError(errors.New("must at least pass one updatedBy in UpdatedByNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updated_by NOT IN (?)", updatedBy))
}

// ===== END of query set FlagSnapshotQuerySet

// ===== BEGIN of FlagSnapshot modifiers

// FlagSnapshotDBSchemaField describes database schema field. It requires for method 'Update'
type FlagSnapshotDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f FlagSnapshotDBSchemaField) String() string {
	return string(f)
}

// FlagSnapshotDBSchema stores db field names of FlagSnapshot
var FlagSnapshotDBSchema = struct {
	ID        FlagSnapshotDBSchemaField
	CreatedAt FlagSnapshotDBSchemaField
	UpdatedAt FlagSnapshotDBSchemaField
	DeletedAt FlagSnapshotDBSchemaField
	FlagID    FlagSnapshotDBSchemaField
	UpdatedBy FlagSnapshotDBSchemaField
	Flag      FlagSnapshotDBSchemaField
}{

	ID:        FlagSnapshotDBSchemaField("id"),
	CreatedAt: FlagSnapshotDBSchemaField("created_at"),
	UpdatedAt: FlagSnapshotDBSchemaField("updated_at"),
	DeletedAt: FlagSnapshotDBSchemaField("deleted_at"),
	FlagID:    FlagSnapshotDBSchemaField("flag_id"),
	UpdatedBy: FlagSnapshotDBSchemaField("updated_by"),
	Flag:      FlagSnapshotDBSchemaField("flag"),
}

// Update updates FlagSnapshot fields by primary key
// nolint: dupl
func (o *FlagSnapshot) Update(db *gorm.DB, fields ...FlagSnapshotDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"created_at": o.CreatedAt,
		"updated_at": o.UpdatedAt,
		"deleted_at": o.DeletedAt,
		"flag_id":    o.FlagID,
		"updated_by": o.UpdatedBy,
		"flag":       o.Flag,
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

		return fmt.Errorf("can't update FlagSnapshot %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// FlagSnapshotUpdater is an FlagSnapshot updates manager
type FlagSnapshotUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewFlagSnapshotUpdater creates new FlagSnapshot updater
// nolint: dupl
func NewFlagSnapshotUpdater(db *gorm.DB) FlagSnapshotUpdater {
	return FlagSnapshotUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&FlagSnapshot{}),
	}
}

// ===== END of FlagSnapshot modifiers

// ===== END of all query sets
