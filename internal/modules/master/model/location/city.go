package location

import (
	"backend-app/internal/base/models"
)

type City struct {
	models.BaseModel
	ProvinceID uint32 `gorm:"column:province_id" json:"province_id"`
	City       string `gorm:"column:city;type:varchar(100)" json:"city"`
}

func (City) TableName() string {
	return "cities_m"
}
