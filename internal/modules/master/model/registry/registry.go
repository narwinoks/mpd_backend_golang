package registry

import (
	"backend-app/internal/base/models"
)

type Registry struct {
	models.BaseModel
	Name      string     `gorm:"column:name;type:varchar(100)" json:"name"`
	Path      string     `gorm:"column:path;type:varchar(100)" json:"path"`
	Icon      string     `gorm:"column:icon;type:varchar(50)" json:"icon"`
	HeadID    *uint32    `gorm:"column:head_id" json:"head_id"`
	SortOrder int        `gorm:"column:sort_order;default:0" json:"sort_order"`
	Children  []Registry `gorm:"foreignKey:HeadID;references:ID" json:"children"`
}

func (Registry) TableName() string {
	return "master_registries_m"
}
