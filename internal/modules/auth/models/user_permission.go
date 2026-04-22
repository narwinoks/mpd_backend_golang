package models

import (
	"backend-app/internal/base/models"
)

type UserPermission struct {
	models.BaseModel
	UserID       uint32 `gorm:"column:user_id" json:"user_id"`
	PermissionID uint32 `gorm:"column:permission_id" json:"permission_id"`
}

func (UserPermission) TableName() string {
	return "user_permission_m"
}
