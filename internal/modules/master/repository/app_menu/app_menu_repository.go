package app_menu

import (
	"backend-app/internal/modules/auth/models"
	req "backend-app/internal/modules/master/request/app_menu"
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

func (r *appMenuRepositoryImpl) FindAll(ctx context.Context, request req.AppMenuFilterRequest) ([]models.AppMenu, int64, error) {
	var results []AppMenuWithCount
	var items []models.AppMenu
	var total int64

	db := r.db.WithContext(ctx).Model(&models.AppMenu{}).
		Preload("AppModule").
		Preload("Parent")

	if request.AppModuleID != "" {
		// Use subquery instead of JOIN to avoid column collision with SELECT * from PaginateScope
		db = db.Preload("SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, app_module_id, parent_id, code, name, path, icon, description, sort_order, is_active, created_at, updated_at").
				Order("sort_order ASC")
		}).
			Preload("SubMenus.AppModule").
			Preload("SubMenus.SubMenus", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, uuid, app_module_id, parent_id, code, name, path, icon, description, sort_order, is_active, created_at, updated_at").
					Order("sort_order ASC")
			}).
			Preload("SubMenus.SubMenus.AppModule")

		db = db.Where("app_menus_m.app_module_id IN (SELECT id FROM app_modules_m WHERE uuid = ?)", request.AppModuleID)
	}

	if request.HeadID != "" {
		db = db.Where("app_menus_m.parent_id IN (SELECT id FROM app_menus_m WHERE uuid = ?)", request.HeadID)
	} else if request.Search == "" {
		// Only filter by root if module is specified and no search is active
		if request.AppModuleID != "" {
			db = db.Where("app_menus_m.parent_id IS NULL")
		}
	}

	err := db.Scopes(pagination.PaginateScope(request.BaseRequest)).
		Scopes(request.SearchScope("app_menus_m.name", "app_menus_m.code")).
		Order("app_menus_m.sort_order ASC").
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
		Preload("SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, app_module_id, parent_id, code, name, path, icon, description, sort_order, is_active, created_at, updated_at").
				Order("sort_order ASC")
		}).
		Preload("SubMenus.AppModule").
		Preload("SubMenus.SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, app_module_id, parent_id, code, name, path, icon, description, sort_order, is_active, created_at, updated_at").
				Order("sort_order ASC")
		}).
		Preload("SubMenus.SubMenus.AppModule").
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
		Preload("SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, app_module_id, parent_id, code, name, path, icon, description, sort_order, is_active, created_at, updated_at").
				Order("sort_order ASC")
		}).
		Preload("SubMenus.AppModule").
		Preload("SubMenus.SubMenus", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, app_module_id, parent_id, code, name, path, icon, description, sort_order, is_active, created_at, updated_at").
				Order("sort_order ASC")
		}).
		Preload("SubMenus.SubMenus.AppModule").
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
