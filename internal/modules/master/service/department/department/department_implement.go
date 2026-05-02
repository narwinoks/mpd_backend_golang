package department

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/department"
	repo "backend-app/internal/modules/master/repository/department/department"
	req "backend-app/internal/modules/master/request/department/department"
	res "backend-app/internal/modules/master/response/department/department"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type departmentServiceImpl struct{ repo repo.DepartmentRepository }

func NewDepartmentService(repo repo.DepartmentRepository) DepartmentService {
	return &departmentServiceImpl{repo: repo}
}

func toResponse(m *model.Department) *res.DepartmentResponse {
	return &res.DepartmentResponse{
		ID:             m.UUID,
		DepartmentName: m.DepartmentName,
		ExternalCode:   m.ExternalCode,
		IsActive:       m.IsActive,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

func (s *departmentServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.DepartmentResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch departments: %v", err)
		return nil, nil, err
	}
	var responses []res.DepartmentResponse
	for _, item := range items {
		responses = append(responses, *toResponse(&item))
	}
	return responses, pagination.BuildMeta(total, request.Page, request.Paginate, len(responses)), nil
}

func (s *departmentServiceImpl) GetByID(ctx context.Context, id string) (*res.DepartmentResponse, error) {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return nil, exception.NewNotFoundError("Data not found")
	}
	return toResponse(m), nil
}

func (s *departmentServiceImpl) Create(ctx context.Context, request req.CreateDepartmentRequest) (*res.DepartmentResponse, error) {
	m := &model.Department{DepartmentName: request.DepartmentName}
	m.ExternalCode = request.ExternalCode
	if err := s.repo.Create(ctx, m); err != nil {
		logrus.Errorf("Failed to create department: %v", err)
		return nil, err
	}
	return toResponse(m), nil
}

func (s *departmentServiceImpl) Update(ctx context.Context, id string, request req.UpdateDepartmentRequest) (*res.DepartmentResponse, error) {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return nil, exception.NewNotFoundError("Data not found")
	}
	m.DepartmentName = request.DepartmentName
	m.ExternalCode = request.ExternalCode
	if request.IsActive != nil {
		m.IsActive = *request.IsActive
	}
	if err := s.repo.Update(ctx, m); err != nil {
		logrus.Errorf("Failed to update department: %v", err)
		return nil, err
	}
	return toResponse(m), nil
}

func (s *departmentServiceImpl) Delete(ctx context.Context, id string) error {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return exception.NewNotFoundError("Data not found")
	}
	if err := s.repo.Delete(ctx, m.ID); err != nil {
		logrus.Errorf("Failed to delete department: %v", err)
		return err
	}
	return nil
}
