package models

import (
	"backend-app/internal/base/models"
	"backend-app/internal/modules/master/model/general"
	"time"
)

type Employee struct {
	models.BaseModel
	ReligionID         uint32           `gorm:"column:religion_id" json:"religion_id"`
	GenderID           uint32           `gorm:"column:gender_id" json:"gender_id"`
	JobTitleID         uint32           `gorm:"column:job_title_id" json:"job_title_id"`
	EmploymentStatusID uint32           `gorm:"column:employment_status_id" json:"employment_status_id"`
	FullName           string           `gorm:"column:full_name;type:varchar(100)" json:"full_name"`
	IdentityNumber     string           `gorm:"column:identity_number;type:varchar(20)" json:"identity_number"`
	NIP                string           `gorm:"column:nip;type:varchar(20)" json:"nip"`
	NPWP               string           `gorm:"column:npwp;type:varchar(20)" json:"npwp"`
	BirthPlace         string           `gorm:"column:birth_place;type:varchar(100)" json:"birth_place"`
	BirthDate          time.Time        `gorm:"column:birth_date;type:date" json:"birth_date"`
	Religion           general.Religion `gorm:"foreignKey:ReligionID;references:ID"`
}

func (Employee) TableName() string {
	return "employees_m"
}
