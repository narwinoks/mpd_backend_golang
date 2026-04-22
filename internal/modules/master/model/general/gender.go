package general

import (
	"backend-app/internal/base/models"
)

type Gender struct {
	models.BaseModel
	Gender string `gorm:"column:gender;type:varchar(20)" json:"gender"`
}

func (Gender) TableName() string {
	return "genders_m"
}
