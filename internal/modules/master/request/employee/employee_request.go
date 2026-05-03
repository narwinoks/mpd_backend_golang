package employee

import (
	"backend-app/pkg/utils"
)

type EmployeeDetailRequest struct {
	MaritalStatusID      string          `json:"marital_status_id" binding:"required"`
	FunctionalPositionID string          `json:"functional_position_id" binding:"required"`
	StructuralPositionID string          `json:"structural_position_id" binding:"required"`
	BankID               *string         `json:"bank_id"`
	BankAccountNumber    string          `json:"bank_account_number" binding:"max=50"`
	BankAccountName      string          `json:"bank_account_name" binding:"max=100"`
	JoinDate             utils.DateOnly  `json:"join_date" binding:"required"`
	ResignDate           *utils.DateOnly `json:"resign_date"`
	RetirementDate       *utils.DateOnly `json:"retirement_date"`
}

type EmployeeAddressRequest struct {
	AddressType   string `json:"address_type" binding:"required,oneof=KTP DOMICILE OFFICE OTHER"`
	FullAddress   string `json:"full_address" binding:"required"`
	ProvinceID    string `json:"province_id" binding:"required"`
	CityID        string `json:"city_id" binding:"required"`
	SubdistrictID string `json:"subdistrict_id" binding:"required"`
	VillageID     string `json:"village_id" binding:"required"`
}

type EmployeeEducationRequest struct {
	EducationLevelID   string         `json:"education_level_id" binding:"required"`
	InstitutionName    string         `json:"institution_name" binding:"required,max=200"`
	InstitutionAddress string         `json:"institution_address" binding:"max=250"`
	Major              string         `json:"major" binding:"required,max=255"`
	StartDate          utils.DateOnly `json:"start_date" binding:"required"`
	GraduationDate     utils.DateOnly `json:"graduation_date" binding:"required"`
	CertificateDate    utils.DateOnly `json:"certificate_date" binding:"required"`
	CertificateNumber  string         `json:"certificate_number" binding:"required,max=255"`
	GPA                string         `json:"gpa" binding:"required,max:2"`
	FrontTitle         string         `json:"front_title" binding:"max=30"`
	BackTitle          string         `json:"back_title" binding:"max=30"`
	IsHighest          bool           `json:"is_highest"`
}

type CreateEmployeeRequest struct {
	ReligionID         string `json:"religion_id" binding:"required"`
	GenderID           string `json:"gender_id" binding:"required"`
	JobTitleID         string `json:"job_title_id" binding:"required"`
	EmploymentStatusID string `json:"employement_status_id" binding:"required"`
	FullName           string `json:"full_name" binding:"required,max=100"`
	IdentityNumber     string `json:"identity_number" binding:"max=20"`
	NIP                string `json:"nip" binding:"max=20"`

	NPWP              string                     `json:"npwp" binding:"max=20,is_npwp"`
	BirthPlace        string                     `json:"birth_place" binding:"required,max=100"`
	BirthDate         utils.DateOnly             `json:"birth_date" binding:"required"`
	EmployeeDetail    EmployeeDetailRequest      `json:"employee_detail" binding:"required"`
	EmployeeAddresses []EmployeeAddressRequest   `json:"employee_addresses" binding:"required,min=1"`
	EmployeeEducation []EmployeeEducationRequest `json:"employee_education" binding:"required,min=1"`
}

type UpdateEmployeeRequest struct {
	ReligionID         string `json:"religion_id" binding:"required"`
	GenderID           string `json:"gender_id" binding:"required"`
	JobTitleID         string `json:"job_title_id" binding:"required"`
	EmploymentStatusID string `json:"employement_status_id" binding:"required"`
	FullName           string `json:"full_name" binding:"required,max=100"`
	IdentityNumber     string `json:"identity_number" binding:"max=20"`
	NIP                string `json:"nip" binding:"max=20"`

	NPWP              string                     `json:"npwp" binding:"max=20"`
	BirthPlace        string                     `json:"birth_place" binding:"required,max=100"`
	BirthDate         utils.DateOnly             `json:"birth_date" binding:"required"`
	EmployeeDetail    EmployeeDetailRequest      `json:"employee_detail" binding:"required"`
	EmployeeAddresses []EmployeeAddressRequest   `json:"employee_addresses" binding:"required,min=1"`
	EmployeeEducation []EmployeeEducationRequest `json:"employee_education" binding:"required,min=1"`
}
