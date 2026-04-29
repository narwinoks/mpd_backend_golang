package registry

import (
	"backend-app/internal/modules/master/model/registry"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type registryRepositoryImpl struct {
	db *gorm.DB
}

func NewRegistryRepository(db *gorm.DB) RegistryRepository {
	return &registryRepositoryImpl{db: db}
}

type RegistryWithCount struct {
	registry.Registry
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *registryRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]registry.Registry, int64, error) {
	var results []RegistryWithCount
	var registries []registry.Registry
	var total int64

	err := r.db.WithContext(ctx).Model(&registry.Registry{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("name", "path")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			registries = append(registries, res.Registry)
		}
	}

	return registries, total, nil
}

func (r *registryRepositoryImpl) FindByID(ctx context.Context, id uint32) (*registry.Registry, error) {
	var reg registry.Registry
	err := r.db.WithContext(ctx).First(&reg, id).Error
	if err != nil {
		return nil, err
	}
	return &reg, nil
}

func (r *registryRepositoryImpl) Create(ctx context.Context, reg *registry.Registry) error {
	return r.db.WithContext(ctx).Create(reg).Error
}

func (r *registryRepositoryImpl) Update(ctx context.Context, reg *registry.Registry) error {
	return r.db.WithContext(ctx).Save(reg).Error
}

func (r *registryRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var reg registry.Registry
	if err := r.db.WithContext(ctx).First(&reg, id).Error; err != nil {
		return err
	}
	return reg.SetNonActive(r.db.WithContext(ctx).Model(&reg))
}

func (r *registryRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*registry.Registry, error) {
	var reg registry.Registry
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&reg).Error
	if err != nil {
		return nil, err
	}
	return &reg, nil
}

func (r *registryRepositoryImpl) FindNested(ctx context.Context) ([]registry.Registry, error) {
	var registries []registry.Registry
	err := r.db.WithContext(ctx).
		Where("head_id IS NULL").
		Order("sort_order ASC").
		Preload("Children", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		}).
		Preload("Children.Children", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		}).
		Find(&registries).Error

	return registries, err
}
