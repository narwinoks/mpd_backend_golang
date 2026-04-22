package models

import (
	"backend-app/internal/base/models"
)

type AppModule struct {
	models.BaseModel
	Code      string `gorm:"column:code;type:varchar(20)" json:"code"`
	Name      string `gorm:"column:name;type:varchar(100)" json:"name"`
	Category  string `gorm:"column:category;type:varchar(50)" json:"category"`
	SortOrder int    `gorm:"column:sort_order" json:"sort_order"`
}

func (AppModule) TableName() string {
	return "app_modules_m"
}
