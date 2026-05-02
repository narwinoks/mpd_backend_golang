package department

import "backend-app/internal/base/models"

type Department struct {
	models.BaseModel
	DepartmentName string `gorm:"column:department_name;type:varchar(255);not null" json:"department_name"`
}

func (Department) TableName() string {
	return "departments_m"
}
