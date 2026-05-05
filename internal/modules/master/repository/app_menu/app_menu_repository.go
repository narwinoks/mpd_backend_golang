package app_menu

import (
	"backend-app/internal/modules/auth/models"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type appMenuRepositoryImpl struct {
	db *gorm.DB
}

func NewAppMenuRepository(db *gorm.DB) AppMenuRepository {
	return &appMenuRepositoryImpl{db: db}
}

type AppMenuWithCount struct {
	models.AppMenu
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *appMenuRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]models.AppMenu, int64, error) {
	var results []AppMenuWithCount
	var items []models.AppMenu
	var total int64

	err := r.db.WithContext(ctx).Model(&models.AppMenu{}).
		Preload("AppModule").
		Preload("Parent").
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("name", "code")).
		Order("sort_order ASC").
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			items = append(items, res.AppMenu)
		}
	}

	return items, total, nil
}

func (r *appMenuRepositoryImpl) FindByID(ctx context.Context, id uint32) (*models.AppMenu, error) {
	var item models.AppMenu
	err := r.db.WithContext(ctx).
		Preload("AppModule").
		Preload("Parent").
		First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *appMenuRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*models.AppMenu, error) {
	var item models.AppMenu
	err := r.db.WithContext(ctx).
		Preload("AppModule").
		Preload("Parent").
		Where("uuid = ?", uuid).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *appMenuRepositoryImpl) Create(ctx context.Context, item *models.AppMenu) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *appMenuRepositoryImpl) Update(ctx context.Context, item *models.AppMenu) error {
	return r.db.WithContext(ctx).Updates(item).Error
}

func (r *appMenuRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var item models.AppMenu
	if err := r.db.WithContext(ctx).First(&item, id).Error; err != nil {
		return err
	}
	return item.SetNonActive(r.db.WithContext(ctx).Model(&item))
}
