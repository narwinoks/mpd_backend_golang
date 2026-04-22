package job

import (
	"backend-app/internal/base/models"
)

type EmploymentStatus struct {
	models.BaseModel
	Code           string `gorm:"column:code;type:varchar(20)" json:"code"`
	EmployeeStatus string `gorm:"column:employee_status;type:varchar(100)" json:"employee_status"`
}

func (EmploymentStatus) TableName() string {
	return "employment_statuses_m"
}
