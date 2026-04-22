package location

import (
	"backend-app/internal/base/models"
)

type Province struct {
	models.BaseModel
	Province string `gorm:"column:province;type:varchar(100)" json:"province"`
}

func (Province) TableName() string {
	return "provinces_m"
}
