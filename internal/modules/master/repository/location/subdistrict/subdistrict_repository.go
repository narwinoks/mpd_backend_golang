package subdistrict

import (
	"backend-app/internal/modules/master/model/location"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type subdistrictRepositoryImpl struct {
	db *gorm.DB
}

func NewSubdistrictRepository(db *gorm.DB) SubdistrictRepository {
	return &subdistrictRepositoryImpl{db: db}
}

type SubdistrictWithCount struct {
	location.Subdistrict
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *subdistrictRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]location.Subdistrict, int64, error) {
	var results []SubdistrictWithCount
	var items []location.Subdistrict
	var total int64

	err := r.db.WithContext(ctx).Model(&location.Subdistrict{}).
		Preload("Province", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "province")
		}).
		Preload("City", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "city")
		}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("subdistrict")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			items = append(items, res.Subdistrict)
		}
	}

	return items, total, nil
}

func (r *subdistrictRepositoryImpl) FindByID(ctx context.Context, id uint32) (*location.Subdistrict, error) {
	var item location.Subdistrict
	err := r.db.WithContext(ctx).
		Preload("Province", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "province")
		}).
		Preload("City", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "city")
		}).
		First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *subdistrictRepositoryImpl) Create(ctx context.Context, subdistrict *location.Subdistrict) error {
	return r.db.WithContext(ctx).Create(subdistrict).Error
}

func (r *subdistrictRepositoryImpl) Update(ctx context.Context, subdistrict *location.Subdistrict) error {
	return r.db.WithContext(ctx).Updates(subdistrict).Error
}

func (r *subdistrictRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var item location.Subdistrict
	if err := r.db.WithContext(ctx).First(&item, id).Error; err != nil {
		return err
	}
	return item.SetNonActive(r.db.WithContext(ctx).Model(&item))
}

func (r *subdistrictRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*location.Subdistrict, error) {
	var item location.Subdistrict
	err := r.db.WithContext(ctx).
		Preload("Province", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "province")
		}).
		Preload("City", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "city")
		}).
		Where("uuid = ?", Uuid).
		First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
