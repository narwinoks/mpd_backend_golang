package models

import (
	"backend-app/internal/base/models"
	"time"
)

type EmployeeEducation struct {
	models.BaseModel
	EmployeeID         uint32    `gorm:"column:employee_id" json:"employee_id"`
	EducationLevelID   uint32    `gorm:"column:education_level_id" json:"education_level_id"`
	InstitutionName    string    `gorm:"column:institution_name;type:varchar(200)" json:"institution_name"`
	InstitutionAddress string    `gorm:"column:institution_address;type:varchar(250)" json:"institution_address"`
	Major              string    `gorm:"column:major;type:varchar(255)" json:"major"`
	StartDate          time.Time `gorm:"column:start_date;type:date" json:"start_date"`
	GraduationDate     time.Time `gorm:"column:graduation_date;type:date" json:"graduation_date"`
	CertificateDate    time.Time `gorm:"column:certificate_date;type:date" json:"certificate_date"`
	CertificateNumber  string    `gorm:"column:certificate_number;type:varchar(255)" json:"certificate_number"`
	GPA                float64   `gorm:"column:gpa;type:decimal(10,2)" json:"gpa"`
	FrontTitle         string    `gorm:"column:front_title;type:varchar(30)" json:"front_title"`
	BackTitle          string    `gorm:"column:back_title;type:varchar(30)" json:"back_title"`
	IsHighest          bool      `gorm:"column:is_highest;default:false" json:"is_highest"`
}

func (EmployeeEducation) TableName() string {
	return "employee_educations_m"
}
