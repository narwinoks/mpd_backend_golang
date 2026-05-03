package general

import (
	"backend-app/internal/base/models"
)

type Education struct {
	models.BaseModel
	EducationType string `gorm:"column:education_type;type:varchar(20)" json:"education_type"`
	Code          string `gorm:"column:code;type:varchar(50)" json:"code"`
	Name          string `gorm:"column:name;type:varchar(100)" json:"name"`
	SortOrder     int    `gorm:"column:sort_order;default:0" json:"sort_order"`
}

func (Education) TableName() string {
	return "educations_m"
}
