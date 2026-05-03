package models

import (
	"backend-app/internal/base/models"
	"time"
)

type EmployeeDetail struct {
	models.BaseModel
	EmployeeID           *uint32    `gorm:"column:employee_id;uniqueIndex" json:"employee_id"`
	MaritalStatusID      uint32     `gorm:"column:marital_status_id" json:"marital_status_id"`
	FunctionalPositionID uint32     `gorm:"column:functional_position_id" json:"functional_position_id"`
	StructuralPositionID uint32     `gorm:"column:structural_position_id" json:"structural_position_id"`
	BankID               *uint32    `gorm:"column:bank_id" json:"bank_id"`
	BankAccountNumber    string     `gorm:"column:bank_account_number;type:varchar(50)" json:"bank_account_number"`
	BankAccountName      string     `gorm:"column:bank_account_name;type:varchar(100)" json:"bank_account_name"`
	JoinDate             time.Time  `gorm:"column:join_date;type:date" json:"join_date"`
	ResignDate           *time.Time `gorm:"column:resign_date;type:date" json:"resign_date"`
	RetirementDate       *time.Time `gorm:"column:retirement_date;type:date" json:"retirement_date"`
}

func (EmployeeDetail) TableName() string {
	return "employee_details_m"
}
