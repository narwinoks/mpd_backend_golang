package app_module

import (
	"backend-app/internal/modules/auth/models"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type appModuleRepositoryImpl struct {
	db *gorm.DB
}

func NewAppModuleRepository(db *gorm.DB) AppModuleRepository {
	return &appModuleRepositoryImpl{db: db}
}

type AppModuleWithCount struct {
	models.AppModule
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *appModuleRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]models.AppModule, int64, error) {
	var results []AppModuleWithCount
	var items []models.AppModule
	var total int64

	err := r.db.WithContext(ctx).Model(&models.AppModule{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("name")).
		Order("sort_order ASC").
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			items = append(items, res.AppModule)
		}
	}

	return items, total, nil
}

func (r *appModuleRepositoryImpl) FindByID(ctx context.Context, id uint32) (*models.AppModule, error) {
	var item models.AppModule
	err := r.db.WithContext(ctx).Select("id", "uuid", "code", "name", "category", "sort_order", "is_active", "created_at", "updated_at").First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *appModuleRepositoryImpl) Create(ctx context.Context, appModule *models.AppModule) error {
	return r.db.WithContext(ctx).Create(appModule).Error
}

func (r *appModuleRepositoryImpl) Update(ctx context.Context, appModule *models.AppModule) error {
	return r.db.WithContext(ctx).Updates(appModule).Error
}

func (r *appModuleRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var item models.AppModule
	if err := r.db.WithContext(ctx).First(&item, id).Error; err != nil {
		return err
	}
	return item.SetNonActive(r.db.WithContext(ctx).Model(&item))
}

func (r *appModuleRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*models.AppModule, error) {
	var item models.AppModule
	err := r.db.WithContext(ctx).Select("id", "uuid", "code", "name", "category", "sort_order", "is_active", "created_at", "updated_at").Where("uuid = ?", uuid).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
