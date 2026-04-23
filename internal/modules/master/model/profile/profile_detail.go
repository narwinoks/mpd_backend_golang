package profile

import (
	"backend-app/internal/base/models"
	"time"
)

type ProfileDetail struct {
	models.BaseModel
	Website          string    `gorm:"column:website" json:"website"`
	Longitude        float64   `gorm:"column:longitude" json:"longitude"`
	Latitude         float64   `gorm:"column:latitude" json:"latitude"`
	RegistrationDate time.Time `gorm:"column:registration_date" json:"registration_date"`
	Moto             string    `gorm:"column:moto" json:"moto"`
}

func (ProfileDetail) TableName() string {
	return "profile_detail_m"
}
