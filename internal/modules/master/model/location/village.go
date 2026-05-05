package location

import (
	"backend-app/internal/base/models"
)

type Village struct {
	models.BaseModel
	ProvinceID    uint32      `gorm:"column:province_id" json:"province_id"`
	CityID        uint32      `gorm:"column:city_id" json:"city_id"`
	SubdistrictID uint32      `gorm:"column:subdistrict_id" json:"subdistrict_id"`
	Code          string      `gorm:"column:code;type:varchar(40)" json:"code"`
	Village       string      `gorm:"column:village;type:varchar(100)" json:"village"`
	PostalCode    string      `gorm:"column:postal_code;type:varchar(10)" json:"postal_code"`
	Longitude     string      `gorm:"column:longitude;type:varchar(50)" json:"longitude"`
	Latitude      string      `gorm:"column:latitude;type:varchar(50)" json:"latitude"`
	Province      Province    `gorm:"foreignKey:ProvinceID;references:ID"`
	City          City        `gorm:"foreignKey:CityID;references:ID"`
	Subdistrict   Subdistrict `gorm:"foreignKey:SubdistrictID;references:ID"`
}

func (Village) TableName() string {
	return "villages_m"
}
