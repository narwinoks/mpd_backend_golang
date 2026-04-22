package models

import (
	"backend-app/internal/base/models"
)

type Role struct {
	models.BaseModel
	Role string `gorm:"column:role;type:varchar(100)" json:"role"`
}

func (Role) TableName() string {
	return "roles_m"
}
