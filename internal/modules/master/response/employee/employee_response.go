package employee

import (
	"backend-app/pkg/utils"
	"time"
)

type GenderResponse struct {
	ID     string `json:"id"`
	Gender string `json:"gender"`
}

type JobTitleResponse struct {
	ID       string `json:"id"`
	JobTitle string `json:"job_title"`
}

type EmploymentStatusResponse struct {
	ID               string `json:"id"`
	EmploymentStatus string `json:"employment_status"`
}

type EmployeeDetailResponse struct {
	MaritalStatusID      uint32          `json:"marital_status_id"`
	FunctionalPositionID uint32          `json:"functional_position_id"`
	StructuralPositionID uint32          `json:"structural_position_id"`
	BankID               *uint32         `json:"bank_id"`
	BankAccountNumber    string          `json:"bank_account_number"`
	BankAccountName      string          `json:"bank_account_name"`
	JoinDate             utils.DateOnly  `json:"join_date"`
	ResignDate           *utils.DateOnly `json:"resign_date"`
	RetirementDate       *utils.DateOnly `json:"retirement_date"`
}

type EmployeeAddressResponse struct {
	ID            string `json:"id"`
	AddressType   string `json:"address_type"`
	FullAddress   string `json:"full_address"`
	ProvinceID    uint32 `json:"province_id"`
	CityID        uint32 `json:"city_id"`
	SubdistrictID uint32 `json:"subdistrict_id"`
	VillageID     uint32 `json:"village_id"`
}

type EmployeeEducationResponse struct {
	ID                 string         `json:"id"`
	EducationLevelID   uint32         `json:"education_level_id"`
	InstitutionName    string         `json:"institution_name"`
	InstitutionAddress string         `json:"institution_address"`
	Major              string         `json:"major"`
	StartDate          utils.DateOnly `json:"start_date"`
	GraduationDate     utils.DateOnly `json:"graduation_date"`
	CertificateDate    utils.DateOnly `json:"certificate_date"`
	CertificateNumber  string         `json:"certificate_number"`
	GPA                string         `json:"gpa"`
	FrontTitle         string         `json:"front_title"`
	BackTitle          string         `json:"back_title"`
	IsHighest          bool           `json:"is_highest"`
}

type EmployeeResponse struct {
	ID                 string                      `json:"id"`
	ReligionID         uint32                      `json:"religion_id"`
	GenderID           uint32                      `json:"gender_id"`
	JobTitleID         uint32                      `json:"job_title_id"`
	EmploymentStatusID uint32                      `json:"employment_status_id"`
	FullName           string                      `json:"full_name"`
	IdentityNumber     string                      `json:"identity_number"`
	NIP                string                      `json:"nip"`
	NPWP               string                      `json:"npwp"`
	BirthPlace         string                      `json:"birth_place"`
	BirthDate          utils.DateOnly              `json:"birth_date"`
	Gender             *GenderResponse             `json:"gender,omitempty"`
	JobTitle           *JobTitleResponse           `json:"job_title,omitempty"`
	EmploymentStatus   *EmploymentStatusResponse   `json:"employment_status,omitempty"`
	Detail             EmployeeDetailResponse      `json:"detail"`
	Addresses          []EmployeeAddressResponse   `json:"addresses"`
	Educations         []EmployeeEducationResponse `json:"educations"`
	CreatedAt          time.Time                   `json:"created_at"`
	UpdatedAt          time.Time                   `json:"updated_at"`
}

type EmployeeListResponse struct {
	ID                 string                    `json:"id"`
	ReligionID         uint32                    `json:"religion_id"`
	GenderID           uint32                    `json:"gender_id"`
	JobTitleID         uint32                    `json:"job_title_id"`
	EmploymentStatusID uint32                    `json:"employment_status_id"`
	FullName           string                    `json:"full_name"`
	IdentityNumber     string                    `json:"identity_number"`
	NIP                string                    `json:"nip"`
	NPWP               string                    `json:"npwp"`
	BirthPlace         string                    `json:"birth_place"`
	BirthDate          utils.DateOnly            `json:"birth_date"`
	Gender             *GenderResponse           `json:"gender,omitempty"`
	JobTitle           *JobTitleResponse         `json:"job_title,omitempty"`
	EmploymentStatus   *EmploymentStatusResponse `json:"employment_status,omitempty"`
	CreatedAt          time.Time                 `json:"created_at"`
	UpdatedAt          time.Time                 `json:"updated_at"`
}
