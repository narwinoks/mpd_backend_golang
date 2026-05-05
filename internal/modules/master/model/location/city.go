package location

import (
	"backend-app/internal/base/models"
)

type City struct {
	models.BaseModel
	ProvinceID uint32   `gorm:"column:province_id" json:"province_id"`
	Code       string   `gorm:"column:code;type:varchar(40)" json:"code"`
	City       string   `gorm:"column:city;type:varchar(100)" json:"city"`
	Province   Province `gorm:"foreignKey:ProvinceID;references:ID"`
}

func (City) TableName() string {
	return "cities_m"
}
