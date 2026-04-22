package job

import (
	"backend-app/internal/base/models"
)

type JobCategory struct {
	models.BaseModel
	Code        string `gorm:"column:code;type:varchar(20)" json:"code"`
	JobCategory string `gorm:"column:job_category;type:varchar(100)" json:"job_category"`
}

func (JobCategory) TableName() string {
	return "job_categories_m"
}
