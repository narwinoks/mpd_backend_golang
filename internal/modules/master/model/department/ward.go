package department

import (
	"backend-app/internal/base/models"
)

type Ward struct {
	models.BaseModel
	WardName          string     `gorm:"column:ward_name;type:varchar(255);not null"  json:"ward_name"`
	DepartmentID      *uint32    `gorm:"column:department_id"                         json:"department_id"`
	IsExecutive       bool       `gorm:"column:is_executive;default:false"            json:"is_executive"`
	Icon              *string    `gorm:"column:icon;type:varchar(100)"                json:"icon"`
	QueueNumberPrefix *string    `gorm:"column:queue_number_prefix;type:varchar(200)" json:"queue_number_prefix"`
	Department        Department `gorm:"foreignKey:DepartmentID;references:ID"`
}

func (Ward) TableName() string {
	return "wards_m"
}
