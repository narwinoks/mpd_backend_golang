package ward

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/department"
	repo "backend-app/internal/modules/master/repository/department/ward"
	req "backend-app/internal/modules/master/request/department/ward"
	res "backend-app/internal/modules/master/response/department/ward"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type wardServiceImpl struct{ repo repo.WardRepository }

func NewWardService(repo repo.WardRepository) WardService {
	return &wardServiceImpl{repo: repo}
}

func (s *wardServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.WardResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch wards: %v", err)
		return nil, nil, err
	}
	responses := res.FromWards(items)
	return responses, pagination.BuildMeta(total, request.Page, request.Paginate, len(responses)), nil
}

func (s *wardServiceImpl) GetByID(ctx context.Context, id string) (*res.WardResponse, error) {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return nil, exception.NewNotFoundError("Data not found")
	}
	return res.FromWard(m), nil
}

func (s *wardServiceImpl) Create(ctx context.Context, request req.CreateWardRequest) (*res.WardResponse, error) {
	m := &model.Ward{
		WardName:          request.WardName,
		DepartmentID:      request.DepartmentID,
		IsExecutive:       request.IsExecutive,
		Icon:              request.Icon,
		QueueNumberPrefix: request.QueueNumberPrefix,
	}
	m.ExternalCode = request.ExternalCode
	if err := s.repo.Create(ctx, m); err != nil {
		logrus.Errorf("Failed to create ward: %v", err)
		return nil, err
	}
	return res.FromWard(m), nil
}

func (s *wardServiceImpl) Update(ctx context.Context, id string, request req.UpdateWardRequest) (*res.WardResponse, error) {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return nil, exception.NewNotFoundError("Data not found")
	}
	m.WardName = request.WardName
	m.DepartmentID = request.DepartmentID
	m.IsExecutive = request.IsExecutive
	m.Icon = request.Icon
	m.QueueNumberPrefix = request.QueueNumberPrefix
	m.ExternalCode = request.ExternalCode
	if request.IsActive != nil {
		m.IsActive = *request.IsActive
	}
	if err := s.repo.Update(ctx, m); err != nil {
		logrus.Errorf("Failed to update ward: %v", err)
		return nil, err
	}
	return res.FromWard(m), nil
}

func (s *wardServiceImpl) Delete(ctx context.Context, id string) error {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return exception.NewNotFoundError("Data not found")
	}
	if err := s.repo.Delete(ctx, m.ID); err != nil {
		logrus.Errorf("Failed to delete ward: %v", err)
		return err
	}
	return nil
}
