package module

import (
	"backend-app/internal/modules/auth/models"

	"gorm.io/gorm"
)

type moduleRepository struct {
	db *gorm.DB
}

func NewModuleRepository(db *gorm.DB) ModuleRepository {
	return &moduleRepository{db: db}
}
func (r *moduleRepository) GetUserModules(userID uint32, roleID uint32) ([]models.AppModule, error) {
	var modules []models.AppModule

	err := r.db.Table("app_modules_m").
		Select("app_modules_m.id, app_modules_m.code, app_modules_m.name, app_modules_m.category, app_modules_m.sort_order").
		Joins("LEFT JOIN role_modules_m ON role_modules_m.modules_id = app_modules_m.id").
		Joins("LEFT JOIN user_modules_m ON user_modules_m.modules_id = app_modules_m.id").
		Where("(role_modules_m.role_id = ? AND role_modules_m.deleted_at IS NULL) OR (user_modules_m.user_id = ? AND user_modules_m.deleted_at IS NULL)", roleID, userID).
		Group("app_modules_m.id").
		Order("app_modules_m.sort_order ASC").
		Find(&modules).Error

	if err != nil {
		return nil, err
	}

	return modules, nil
}
