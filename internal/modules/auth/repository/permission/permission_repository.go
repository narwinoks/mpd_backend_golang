package permission

import (
	"gorm.io/gorm"
)

type permissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepositoryImpl{db: db}
}

func (r *permissionRepositoryImpl) GetUserPermissions(userID uint32, roleID uint32) ([]string, error) {
	var permissions []string

	// Query permissions from role_permission_m and user_permission_m
	err := r.db.Table("app_permission_m").
		Select("DISTINCT app_permission_m.permission").
		Joins("LEFT JOIN role_permission_m ON role_permission_m.permission_id = app_permission_m.id").
		Joins("LEFT JOIN user_permission_m ON user_permission_m.permission_id = app_permission_m.id").
		Where("(role_permission_m.role_id = ? AND role_permission_m.deleted_at IS NULL) OR (user_permission_m.user_id = ? AND user_permission_m.deleted_at IS NULL)", roleID, userID).
		Find(&permissions).Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
