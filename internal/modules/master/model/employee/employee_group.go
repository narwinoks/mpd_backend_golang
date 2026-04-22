package models

import (
	"backend-app/internal/base/models"
)

type EmployeeGroup struct {
	models.BaseModel
	EmployeeGroup string `gorm:"column:employee_group;type:varchar(100)" json:"employee_group"`
}

func (EmployeeGroup) TableName() string {
	return "employee_group_m"
}
