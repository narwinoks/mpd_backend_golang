package employee

import (
	"backend-app/internal/core/database"
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/employee"
	repo "backend-app/internal/modules/master/repository/employee"
	req "backend-app/internal/modules/master/request/employee"
	res "backend-app/internal/modules/master/response/employee"
	"backend-app/pkg/pagination"
	"backend-app/pkg/utils"
	"context"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type employeeServiceImpl struct {
	repo repo.EmployeeRepository
	db   *gorm.DB
}

func NewEmployeeService(repo repo.EmployeeRepository, db *gorm.DB) EmployeeService {
	return &employeeServiceImpl{
		repo: repo,
		db:   db,
	}
}

func (s *employeeServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.EmployeeListResponse, *pagination.Meta, error) {
	employees, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch employees: %v", err)
		return nil, nil, err
	}

	var response []res.EmployeeListResponse
	for _, e := range employees {
		item := res.EmployeeListResponse{
			ID:                 e.UUID,
			ReligionID:         e.ReligionID,
			GenderID:           e.GenderID,
			JobTitleID:         e.JobTitleID,
			EmploymentStatusID: e.EmploymentStatusID,
			FullName:           e.FullName,
			IdentityNumber:     e.IdentityNumber,
			NIP:                e.NIP,
			NPWP:               e.NPWP,
			BirthPlace:         e.BirthPlace,
			BirthDate:          utils.DateOnly{Time: e.BirthDate},
			CreatedAt:          e.CreatedAt,
			UpdatedAt:          e.UpdatedAt,
		}

		if e.Gender.ID != 0 {
			item.Gender = &res.GenderResponse{
				ID:     e.Gender.UUID,
				Gender: e.Gender.Gender,
			}
		}

		if e.JobTitle.ID != 0 {
			item.JobTitle = &res.JobTitleResponse{
				ID:       e.JobTitle.UUID,
				JobTitle: e.JobTitle.JobTitle,
			}
		}

		if e.EmployeeStatus.ID != 0 {
			item.EmploymentStatus = &res.EmploymentStatusResponse{
				ID:               e.EmployeeStatus.UUID,
				EmploymentStatus: e.EmployeeStatus.EmployeeStatus,
			}
		}

		response = append(response, item)
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *employeeServiceImpl) GetByID(ctx context.Context, id string) (*res.EmployeeResponse, error) {
	employee, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Employee not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return s.mapToResponse(employee), nil
}

func (s *employeeServiceImpl) Create(ctx context.Context, request req.CreateEmployeeRequest) (string, error) {
	religionID, _ := database.ResolveUUID(ctx, s.db, "religions_m", request.ReligionID)
	genderID, _ := database.ResolveUUID(ctx, s.db, "genders_m", request.GenderID)
	jobTitleID, _ := database.ResolveUUID(ctx, s.db, "job_titles_m", request.JobTitleID)
	employmentStatusID, _ := database.ResolveUUID(ctx, s.db, "employment_statuses_m", request.EmploymentStatusID)

	maritalStatusID, _ := database.ResolveUUID(ctx, s.db, "marital_status_m", request.EmployeeDetail.MaritalStatusID)
	functionalPositionID, _ := database.ResolveUUID(ctx, s.db, "positions_m", request.EmployeeDetail.FunctionalPositionID)
	structuralPositionID, _ := database.ResolveUUID(ctx, s.db, "positions_m", request.EmployeeDetail.StructuralPositionID)
	var bankID *uint32
	if request.EmployeeDetail.BankID != nil {
		id, _ := database.ResolveUUID(ctx, s.db, "banks_m", *request.EmployeeDetail.BankID)
		bankID = &id
	}

	employee := &model.Employee{
		ReligionID:         religionID,
		GenderID:           genderID,
		JobTitleID:         jobTitleID,
		EmploymentStatusID: employmentStatusID,
		FullName:           request.FullName,
		IdentityNumber:     request.IdentityNumber,
		NIP:                request.NIP,
		NPWP:               request.NPWP,
		BirthPlace:         request.BirthPlace,
		BirthDate:          request.BirthDate.Time,
		Detail: model.EmployeeDetail{
			MaritalStatusID:      maritalStatusID,
			FunctionalPositionID: functionalPositionID,
			StructuralPositionID: structuralPositionID,
			BankID:               bankID,
			BankAccountNumber:    request.EmployeeDetail.BankAccountNumber,
			BankAccountName:      request.EmployeeDetail.BankAccountName,
			JoinDate:             request.EmployeeDetail.JoinDate.Time,
		},
	}

	if request.EmployeeDetail.ResignDate != nil {
		t := request.EmployeeDetail.ResignDate.Time
		employee.Detail.ResignDate = &t
	}
	if request.EmployeeDetail.RetirementDate != nil {
		t := request.EmployeeDetail.RetirementDate.Time
		employee.Detail.RetirementDate = &t
	}

	for _, addr := range request.EmployeeAddresses {
		provinceID, _ := database.ResolveUUID(ctx, s.db, "provinces_m", addr.ProvinceID)
		cityID, _ := database.ResolveUUID(ctx, s.db, "cities_m", addr.CityID)
		subdistrictID, _ := database.ResolveUUID(ctx, s.db, "subdistrict_m", addr.SubdistrictID)
		villageID, _ := database.ResolveUUID(ctx, s.db, "villages_m", addr.VillageID)

		employee.Addresses = append(employee.Addresses, model.EmployeeAddress{
			AddressType:   addr.AddressType,
			FullAddress:   addr.FullAddress,
			ProvinceID:    provinceID,
			CityID:        cityID,
			SubdistrictID: subdistrictID,
			VillageID:     villageID,
		})
	}

	for _, edu := range request.EmployeeEducation {
		eduLevelID, _ := database.ResolveUUID(ctx, s.db, "educations_m", edu.EducationLevelID)
		gpa, _ := strconv.ParseFloat(edu.GPA, 64)

		employee.Educations = append(employee.Educations, model.EmployeeEducation{
			EducationLevelID:   eduLevelID,
			InstitutionName:    edu.InstitutionName,
			InstitutionAddress: edu.InstitutionAddress,
			Major:              edu.Major,
			StartDate:          edu.StartDate.Time,
			GraduationDate:     edu.GraduationDate.Time,
			CertificateDate:    edu.CertificateDate.Time,
			CertificateNumber:  edu.CertificateNumber,
			GPA:                gpa,
			FrontTitle:         edu.FrontTitle,
			BackTitle:          edu.BackTitle,
			IsHighest:          edu.IsHighest,
		})
	}

	err := s.repo.Create(ctx, employee)
	if err != nil {
		logrus.Errorf("Failed to create employee: %v", err)
		return "", err
	}

	return employee.UUID, nil
}

func (s *employeeServiceImpl) Update(ctx context.Context, id string, request req.UpdateEmployeeRequest) (string, error) {
	employee, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Employee not found for update with id %s: %v", id, err)
		return "", exception.NewNotFoundError("Data not found")
	}

	religionID, _ := database.ResolveUUID(ctx, s.db, "religions_m", request.ReligionID)
	genderID, _ := database.ResolveUUID(ctx, s.db, "genders_m", request.GenderID)
	jobTitleID, _ := database.ResolveUUID(ctx, s.db, "job_titles_m", request.JobTitleID)
	employmentStatusID, _ := database.ResolveUUID(ctx, s.db, "employment_statuses_m", request.EmploymentStatusID)

	maritalStatusID, _ := database.ResolveUUID(ctx, s.db, "marital_status_m", request.EmployeeDetail.MaritalStatusID)
	functionalPositionID, _ := database.ResolveUUID(ctx, s.db, "positions_m", request.EmployeeDetail.FunctionalPositionID)
	structuralPositionID, _ := database.ResolveUUID(ctx, s.db, "positions_m", request.EmployeeDetail.StructuralPositionID)
	var bankID *uint32
	if request.EmployeeDetail.BankID != nil {
		id, _ := database.ResolveUUID(ctx, s.db, "banks_m", *request.EmployeeDetail.BankID)
		bankID = &id
	}

	employee.ReligionID = religionID
	employee.GenderID = genderID
	employee.JobTitleID = jobTitleID
	employee.EmploymentStatusID = employmentStatusID
	employee.FullName = request.FullName
	employee.IdentityNumber = request.IdentityNumber
	employee.NIP = request.NIP
	employee.NPWP = request.NPWP
	employee.BirthPlace = request.BirthPlace
	employee.BirthDate = request.BirthDate.Time

	employee.Detail.MaritalStatusID = maritalStatusID
	employee.Detail.FunctionalPositionID = functionalPositionID
	employee.Detail.StructuralPositionID = structuralPositionID
	employee.Detail.BankID = bankID
	employee.Detail.BankAccountNumber = request.EmployeeDetail.BankAccountNumber
	employee.Detail.BankAccountName = request.EmployeeDetail.BankAccountName
	employee.Detail.JoinDate = request.EmployeeDetail.JoinDate.Time

	if request.EmployeeDetail.ResignDate != nil {
		t := request.EmployeeDetail.ResignDate.Time
		employee.Detail.ResignDate = &t
	} else {
		employee.Detail.ResignDate = nil
	}

	if request.EmployeeDetail.RetirementDate != nil {
		t := request.EmployeeDetail.RetirementDate.Time
		employee.Detail.RetirementDate = &t
	} else {
		employee.Detail.RetirementDate = nil
	}

	employee.Addresses = nil
	for _, addr := range request.EmployeeAddresses {
		provinceID, _ := database.ResolveUUID(ctx, s.db, "provinces_m", addr.ProvinceID)
		cityID, _ := database.ResolveUUID(ctx, s.db, "cities_m", addr.CityID)
		subdistrictID, _ := database.ResolveUUID(ctx, s.db, "subdistrict_m", addr.SubdistrictID)
		villageID, _ := database.ResolveUUID(ctx, s.db, "villages_m", addr.VillageID)

		employee.Addresses = append(employee.Addresses, model.EmployeeAddress{
			AddressType:   addr.AddressType,
			FullAddress:   addr.FullAddress,
			ProvinceID:    provinceID,
			CityID:        cityID,
			SubdistrictID: subdistrictID,
			VillageID:     villageID,
		})
	}

	employee.Educations = nil
	for _, edu := range request.EmployeeEducation {
		eduLevelID, _ := database.ResolveUUID(ctx, s.db, "educations_m", edu.EducationLevelID)
		gpa, _ := strconv.ParseFloat(edu.GPA, 64)

		employee.Educations = append(employee.Educations, model.EmployeeEducation{
			EducationLevelID:   eduLevelID,
			InstitutionName:    edu.InstitutionName,
			InstitutionAddress: edu.InstitutionAddress,
			Major:              edu.Major,
			StartDate:          edu.StartDate.Time,
			GraduationDate:     edu.GraduationDate.Time,
			CertificateDate:    edu.CertificateDate.Time,
			CertificateNumber:  edu.CertificateNumber,
			GPA:                gpa,
			FrontTitle:         edu.FrontTitle,
			BackTitle:          edu.BackTitle,
			IsHighest:          edu.IsHighest,
		})
	}

	err = s.repo.Update(ctx, employee)
	if err != nil {
		logrus.Errorf("Failed to update employee: %v", err)
		return "", err
	}

	return employee.UUID, nil
}

func (s *employeeServiceImpl) Delete(ctx context.Context, id string) error {
	employee, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Employee not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, employee.ID)
	if err != nil {
		logrus.Errorf("Failed to delete employee: %v", err)
		return err
	}

	return nil
}

func (s *employeeServiceImpl) mapToResponse(e *model.Employee) *res.EmployeeResponse {
	response := &res.EmployeeResponse{
		ID:                 e.UUID,
		ReligionID:         e.ReligionID,
		GenderID:           e.GenderID,
		JobTitleID:         e.JobTitleID,
		EmploymentStatusID: e.EmploymentStatusID,
		FullName:           e.FullName,
		IdentityNumber:     e.IdentityNumber,
		NIP:                e.NIP,
		NPWP:               e.NPWP,
		BirthPlace:         e.BirthPlace,
		BirthDate:          utils.DateOnly{Time: e.BirthDate},
		CreatedAt:          e.CreatedAt,
		UpdatedAt:          e.UpdatedAt,
		Detail: res.EmployeeDetailResponse{
			MaritalStatusID:      e.Detail.MaritalStatusID,
			FunctionalPositionID: e.Detail.FunctionalPositionID,
			StructuralPositionID: e.Detail.StructuralPositionID,
			BankID:               e.Detail.BankID,
			BankAccountNumber:    e.Detail.BankAccountNumber,
			BankAccountName:      e.Detail.BankAccountName,
			JoinDate:             utils.DateOnly{Time: e.Detail.JoinDate},
		},
	}

	if e.Detail.ResignDate != nil {
		response.Detail.ResignDate = &utils.DateOnly{Time: *e.Detail.ResignDate}
	}
	if e.Detail.RetirementDate != nil {
		response.Detail.RetirementDate = &utils.DateOnly{Time: *e.Detail.RetirementDate}
	}

	if e.Gender.ID != 0 {
		response.Gender = &res.GenderResponse{
			ID:     e.Gender.UUID,
			Gender: e.Gender.Gender,
		}
	}

	if e.JobTitle.ID != 0 {
		response.JobTitle = &res.JobTitleResponse{
			ID:       e.JobTitle.UUID,
			JobTitle: e.JobTitle.JobTitle,
		}
	}

	if e.EmployeeStatus.ID != 0 {
		response.EmploymentStatus = &res.EmploymentStatusResponse{
			ID:               e.EmployeeStatus.UUID,
			EmploymentStatus: e.EmployeeStatus.EmployeeStatus,
		}
	}

	for _, addr := range e.Addresses {
		response.Addresses = append(response.Addresses, res.EmployeeAddressResponse{
			ID:            addr.UUID,
			AddressType:   addr.AddressType,
			FullAddress:   addr.FullAddress,
			ProvinceID:    addr.ProvinceID,
			CityID:        addr.CityID,
			SubdistrictID: addr.SubdistrictID,
			VillageID:     addr.VillageID,
		})
	}

	for _, edu := range e.Educations {
		response.Educations = append(response.Educations, res.EmployeeEducationResponse{
			ID:                 edu.UUID,
			EducationLevelID:   edu.EducationLevelID,
			InstitutionName:    edu.InstitutionName,
			InstitutionAddress: edu.InstitutionAddress,
			Major:              edu.Major,
			StartDate:          utils.DateOnly{Time: edu.StartDate},
			GraduationDate:     utils.DateOnly{Time: edu.GraduationDate},
			CertificateDate:    utils.DateOnly{Time: edu.CertificateDate},
			CertificateNumber:  edu.CertificateNumber,
			GPA:                strconv.FormatFloat(edu.GPA, 'f', 2, 64),
			FrontTitle:         edu.FrontTitle,
			BackTitle:          edu.BackTitle,
			IsHighest:          edu.IsHighest,
		})
	}

	return response
}
