package general

import (
	"backend-app/internal/base/models"
)

type Religion struct {
	models.BaseModel
	Religion string `gorm:"column:religion;type:varchar(100)" json:"religion"`
}

func (Religion) TableName() string {
	return "religions_m"
}
