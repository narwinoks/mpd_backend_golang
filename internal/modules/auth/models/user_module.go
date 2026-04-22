package models

import (
	"backend-app/internal/base/models"
)

type UserModule struct {
	models.BaseModel
	UserID    uint32 `gorm:"column:user_id" json:"user_id"`
	ModulesID uint32 `gorm:"column:modules_id" json:"modules_id"`
}

func (UserModule) TableName() string {
	return "user_modules_m"
}
