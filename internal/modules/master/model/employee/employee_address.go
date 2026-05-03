package models

import (
	"backend-app/internal/base/models"
)

type EmployeeAddress struct {
	models.BaseModel
	EmployeeID    uint32 `gorm:"column:employee_id" json:"employee_id"`
	AddressType   string `gorm:"column:address_type;type:varchar(50)" json:"address_type"`
	FullAddress   string `gorm:"column:full_address;type:text" json:"full_address"`
	ProvinceID    uint32 `gorm:"column:province_id" json:"province_id"`
	CityID        uint32 `gorm:"column:city_id" json:"city_id"`
	SubdistrictID uint32 `gorm:"column:subdistrict_id" json:"subdistrict_id"`
	VillageID     uint32 `gorm:"column:village_id" json:"village_id"`
}

func (EmployeeAddress) TableName() string {
	return "employee_addresses_m"
}
