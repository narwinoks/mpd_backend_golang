package location

import (
	"backend-app/internal/base/models"
)

type Subdistrict struct {
	models.BaseModel
	CityID      uint32 `gorm:"column:city_id" json:"city_id"`
	ProvinceID  uint32 `gorm:"column:province_id" json:"province_id"`
	Subdistrict string `gorm:"column:subdistrict;type:varchar(100)" json:"subdistrict"`
}

func (Subdistrict) TableName() string {
	return "subdistrict_m"
}
