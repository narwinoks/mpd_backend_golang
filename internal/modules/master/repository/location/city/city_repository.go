package city

import (
	"backend-app/internal/core/database"
	"backend-app/internal/modules/master/model/location"
	req "backend-app/internal/modules/master/request/location/city"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type cityRepositoryImpl struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityRepositoryImpl{db: db}
}

type CityWithCount struct {
	location.City
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *cityRepositoryImpl) FindAll(ctx context.Context, req req.FindAllRequest) ([]location.City, int64, error) {
	var results []CityWithCount
	var cities []location.City
	var total int64

	query := r.db.WithContext(ctx).Model(&location.City{}).
		Preload("Province")

	if req.ProvinceID != "" {
		provinceID, err := database.ResolveUUID(ctx, r.db, "provinces_m", req.ProvinceID)
		if err == nil {
			query = query.Where("province_id = ?", provinceID)
		}
	}

	err := query.Scopes(pagination.PaginateScope(req.BaseRequest)).
		Scopes(req.SearchScope("city")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			cities = append(cities, res.City)
		}
	}

	return cities, total, nil
}

func (r *cityRepositoryImpl) FindByID(ctx context.Context, id uint32) (*location.City, error) {
	var city location.City
	err := r.db.WithContext(ctx).
		Preload("Province").
		Select("id", "uuid", "province_id", "city", "is_active", "created_at", "updated_at").
		First(&city, id).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *cityRepositoryImpl) Create(ctx context.Context, city *location.City) error {
	return r.db.WithContext(ctx).Create(city).Error
}

func (r *cityRepositoryImpl) Update(ctx context.Context, city *location.City) error {
	return r.db.WithContext(ctx).Updates(city).Error
}

func (r *cityRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var city location.City
	if err := r.db.WithContext(ctx).First(&city, id).Error; err != nil {
		return err
	}
	return city.SetNonActive(r.db.WithContext(ctx).Model(&city))
}

func (r *cityRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*location.City, error) {
	var city location.City
	err := r.db.WithContext(ctx).
		Preload("Province").
		Select("id", "uuid", "province_id", "city", "is_active", "created_at", "updated_at").
		Where("uuid = ?", Uuid).
		First(&city).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}
