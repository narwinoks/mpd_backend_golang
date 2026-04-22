package models

import (
	"backend-app/internal/base/models"
)

type RoleModule struct {
	models.BaseModel
	RoleID    uint32 `gorm:"column:role_id" json:"role_id"`
	ModulesID uint32 `gorm:"column:modules_id" json:"modules_id"`
}

func (RoleModule) TableName() string {
	return "role_modules_m"
}
