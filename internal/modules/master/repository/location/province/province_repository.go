package province

import (
	"backend-app/internal/modules/master/model/location"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type provinceRepositoryImpl struct {
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &provinceRepositoryImpl{db: db}
}

type ProvinceWithCount struct {
	location.Province
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *provinceRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]location.Province, int64, error) {
	var results []ProvinceWithCount
	var items []location.Province
	var total int64

	err := r.db.WithContext(ctx).Model(&location.Province{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("province")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			items = append(items, res.Province)
		}
	}

	return items, total, nil
}

func (r *provinceRepositoryImpl) FindByID(ctx context.Context, id uint32) (*location.Province, error) {
	var province location.Province
	err := r.db.WithContext(ctx).Select("id", "uuid", "province", "is_active", "created_at", "updated_at").First(&province, id).Error
	if err != nil {
		return nil, err
	}
	return &province, nil
}

func (r *provinceRepositoryImpl) Create(ctx context.Context, province *location.Province) error {
	return r.db.WithContext(ctx).Create(province).Error
}

func (r *provinceRepositoryImpl) Update(ctx context.Context, province *location.Province) error {
	return r.db.WithContext(ctx).Updates(province).Error
}

func (r *provinceRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var province location.Province
	if err := r.db.WithContext(ctx).First(&province, id).Error; err != nil {
		return err
	}
	return province.SetNonActive(r.db.WithContext(ctx).Model(&province))
}

func (r *provinceRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*location.Province, error) {
	var province location.Province
	err := r.db.WithContext(ctx).Select("id", "uuid", "province", "is_active", "created_at", "updated_at").Where("uuid = ?", Uuid).First(&province).Error
	if err != nil {
		return nil, err
	}
	return &province, nil
}
