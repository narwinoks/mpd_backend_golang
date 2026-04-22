package profile

import (
	"backend-app/internal/base/models"
)

type Profile struct {
	models.BaseModel
	ProvinceID     uint32 `gorm:"column:province_id" json:"province_id"`
	CityID         uint32 `gorm:"column:city_id" json:"city_id"`
	SubdistrictID  uint32 `gorm:"column:subdistrict_id" json:"subdistrict_id"`
	VillageID      uint32 `gorm:"column:village_id" json:"village_id"`
	PostalCode     string `gorm:"column:postal_code;type:varchar(10)" json:"postal_code"`
	Email          string `gorm:"column:email;type:varchar(100)" json:"email"`
	Name           string `gorm:"column:name;type:varchar(100)" json:"name"`
	Profile        string `gorm:"column:profile;type:text" json:"profile"`
	GovernmentName string `gorm:"column:government_name;type:varchar(100)" json:"government_name"`
	Phone          string `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Telp           string `gorm:"column:telp;type:varchar(20)" json:"telp"`
	FullAddress    string `gorm:"column:full_address;type:text" json:"full_address"`
}

func (Profile) TableName() string {
	return "profiles_m"
}
