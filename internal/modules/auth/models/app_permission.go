package models

import (
	"backend-app/internal/base/models"
)

type AppPermission struct {
	models.BaseModel
	Permission string `gorm:"column:permission;type:varchar(100)" json:"permission"`
}

func (AppPermission) TableName() string {
	return "app_permission_m"
}
