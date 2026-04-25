package menu

import (
	"backend-app/internal/modules/auth/models"

	"gorm.io/gorm"
)

type menuRepositoryImpl struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuInterface {
	return &menuRepositoryImpl{db: db}
}

func (r *menuRepositoryImpl) GetMenusByModuleIDs(moduleIDs []uint32) ([]models.AppMenu, error) {
	var menus []models.AppMenu
	err := r.db.
		Select("id, app_module_id, parent_id, code, name, path, icon, description, sort_order").
		Where("app_module_id IN ?", moduleIDs).
		Order("sort_order ASC").
		Find(&menus).Error

	if err != nil {
		return nil, err
	}
	return menus, nil
}
