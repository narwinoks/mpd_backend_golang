package location

import (
	"backend-app/internal/base/models"
)

type Village struct {
	models.BaseModel
	ProvinceID    uint32 `gorm:"column:province_id" json:"province_id"`
	CityID        uint32 `gorm:"column:city_id" json:"city_id"`
	SubDistrictID uint32 `gorm:"column:subdistrict_id" json:"subdistrict_id"`
	Village       string `gorm:"column:village;type:varchar(100)" json:"city"`
}

func (Village) TableName() string {
	return "villages_m"
}
