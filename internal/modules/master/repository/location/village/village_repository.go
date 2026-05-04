package village

import (
	"backend-app/internal/modules/master/model/location"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type villageRepositoryImpl struct {
	db *gorm.DB
}

func NewVillageRepository(db *gorm.DB) VillageRepository {
	return &villageRepositoryImpl{db: db}
}

type VillageWithCount struct {
	location.Village
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *villageRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]location.Village, int64, error) {
	var results []VillageWithCount
	var items []location.Village
	var total int64

	err := r.db.WithContext(ctx).Model(&location.Village{}).
		Preload("Province", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "province")
		}).
		Preload("City", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "city")
		}).
		Preload("Subdistrict", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "subdistrict")
		}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("village")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			items = append(items, res.Village)
		}
	}

	return items, total, nil
}

func (r *villageRepositoryImpl) FindByID(ctx context.Context, id uint32) (*location.Village, error) {
	var item location.Village
	err := r.db.WithContext(ctx).
		Preload("Province", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "province")
		}).
		Preload("City", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "city")
		}).
		Preload("Subdistrict", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "subdistrict")
		}).
		First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *villageRepositoryImpl) Create(ctx context.Context, village *location.Village) error {
	return r.db.WithContext(ctx).Create(village).Error
}

func (r *villageRepositoryImpl) Update(ctx context.Context, village *location.Village) error {
	return r.db.WithContext(ctx).Updates(village).Error
}

func (r *villageRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var item location.Village
	if err := r.db.WithContext(ctx).First(&item, id).Error; err != nil {
		return err
	}
	return item.SetNonActive(r.db.WithContext(ctx).Model(&item))
}

func (r *villageRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*location.Village, error) {
	var item location.Village
	err := r.db.WithContext(ctx).
		Preload("Province", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "province")
		}).
		Preload("City", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "city")
		}).
		Preload("Subdistrict", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "subdistrict")
		}).
		Where("uuid = ?", Uuid).
		First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
