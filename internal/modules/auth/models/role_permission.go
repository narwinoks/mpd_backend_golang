package models

import (
	"backend-app/internal/base/models"
)

type RolePermission struct {
	models.BaseModel
	RoleID       uint32 `gorm:"column:role_id" json:"role_id"`
	PermissionID uint32 `gorm:"column:permission_id" json:"permission_id"`
}

func (RolePermission) TableName() string {
	return "role_permission_m"
}
