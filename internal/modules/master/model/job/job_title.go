package job

import (
	"backend-app/internal/base/models"
)

type JobTitle struct {
	models.BaseModel
	JobCategoryID uint32      `gorm:"column:job_category_id" json:"job_category_id"`
	JobCategory   JobCategory `gorm:"foreignKey:JobCategoryID;references:ID" json:"job_category"`
	Code          string      `gorm:"column:code;type:varchar(20)" json:"code"`
	JobTitle      string      `gorm:"column:job_title;type:varchar(100)" json:"job_title"`
}

func (JobTitle) TableName() string {
	return "job_titles_m"
}
