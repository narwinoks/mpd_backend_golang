package employment_status

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/master/model/job"
	repo "backend-app/internal/modules/master/repository/employee/employment_status"
	req "backend-app/internal/modules/master/request/employee/employment_status"
	res "backend-app/internal/modules/master/response/employee/employment_status"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type employmentStatusServiceImpl struct {
	repo repo.EmploymentStatusRepository
}

func NewEmploymentStatusService(repo repo.EmploymentStatusRepository) EmploymentStatusService {
	return &employmentStatusServiceImpl{
		repo: repo,
	}
}

func (s *employmentStatusServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.EmploymentStatusResponse, *pagination.Meta, error) {
	entities, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch employment statuses: %v", err)
		return nil, nil, err
	}

	var response []res.EmploymentStatusResponse
	for _, e := range entities {
		response = append(response, res.EmploymentStatusResponse{
			ID:             e.UUID,
			Code:           e.Code,
			EmployeeStatus: e.EmployeeStatus,
			CreatedAt:      e.CreatedAt,
			UpdatedAt:      e.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *employmentStatusServiceImpl) GetByID(ctx context.Context, id string) (*res.EmploymentStatusResponse, error) {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Employment status not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.EmploymentStatusResponse{
		ID:             entity.UUID,
		Code:           entity.Code,
		EmployeeStatus: entity.EmployeeStatus,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}, nil
}

func (s *employmentStatusServiceImpl) Create(ctx context.Context, request req.CreateEmploymentStatusRequest) (*res.EmploymentStatusResponse, error) {
	entity := &job.EmploymentStatus{
		Code:           request.Code,
		EmployeeStatus: request.EmployeeStatus,
	}

	err := s.repo.Create(ctx, entity)
	if err != nil {
		logrus.Errorf("Failed to create employment status: %v", err)
		return nil, err
	}

	return &res.EmploymentStatusResponse{
		ID:             entity.UUID,
		Code:           entity.Code,
		EmployeeStatus: entity.EmployeeStatus,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}, nil
}

func (s *employmentStatusServiceImpl) Update(ctx context.Context, id string, request req.UpdateEmploymentStatusRequest) (*res.EmploymentStatusResponse, error) {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Employment status not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	entity.Code = request.Code
	entity.EmployeeStatus = request.EmployeeStatus

	err = s.repo.Update(ctx, entity)
	if err != nil {
		logrus.Errorf("Failed to update employment status: %v", err)
		return nil, err
	}

	return &res.EmploymentStatusResponse{
		ID:             entity.UUID,
		Code:           entity.Code,
		EmployeeStatus: entity.EmployeeStatus,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}, nil
}

func (s *employmentStatusServiceImpl) Delete(ctx context.Context, id string) error {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Employment status not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, entity.ID)
	if err != nil {
		logrus.Errorf("Failed to delete employment status: %v", err)
		return err
	}

	return nil
}
