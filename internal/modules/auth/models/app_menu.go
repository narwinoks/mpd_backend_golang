package models

import (
	"backend-app/internal/base/models"
)

type AppMenu struct {
	models.BaseModel
	AppModuleID uint32     `gorm:"column:app_module_id" json:"app_module_id"`
	AppModule   AppModule  `gorm:"foreignKey:AppModuleID;references:ID" json:"app_module"`
	ParentID    *uint32    `gorm:"column:parent_id" json:"parent_id"`
	Parent      *AppMenu   `gorm:"foreignKey:ParentID;references:ID" json:"parent"`
	Code        string     `gorm:"column:code;type:varchar(20)" json:"code"`
	Name        string     `gorm:"column:name;type:varchar(100)" json:"name"`
	Path        string     `gorm:"column:path;type:varchar(255)" json:"path"`
	Description string     `gorm:"column:description;type:text" json:"description"`
	Icon        string     `gorm:"column:icon;type:varchar(100)" json:"icon"`
	SortOrder   int        `gorm:"column:sort_order" json:"sort_order"`
	SubMenus    []*AppMenu `gorm:"foreignKey:ParentID" json:"sub_menus"`
}

func (AppMenu) TableName() string {
	return "app_menus_m"
}
