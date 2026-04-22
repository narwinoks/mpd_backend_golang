package models

import (
	"backend-app/internal/base/models"
)

type AppMenu struct {
	models.BaseModel
	AppModuleID uint32  `gorm:"column:app_module_id" json:"app_module_id"`
	ParentID    *uint32 `gorm:"column:parent_id" json:"parent_id"`
	Code        string  `gorm:"column:code;type:varchar(20)" json:"code"`
	Name        string  `gorm:"column:name;type:varchar(100)" json:"name"`
	Path        string  `gorm:"column:path;type:varchar(255)" json:"path"`
	Description string  `gorm:"column:description;type:text" json:"description"`
	Icon        string  `gorm:"column:icon;type:varchar(100)" json:"icon"`
	SortOrder   int     `gorm:"column:sort_order" json:"sort_order"`
}

func (AppMenu) TableName() string {
	return "app_menus_m"
}
