package marital_status

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/general"
	repo "backend-app/internal/modules/master/repository/general/marital_status"
	req "backend-app/internal/modules/master/request/general/marital_status"
	res "backend-app/internal/modules/master/response/general/marital_status"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type maritalStatusServiceImpl struct {
	repo repo.MaritalStatusRepository
}

func NewMaritalStatusService(repo repo.MaritalStatusRepository) MaritalStatusService {
	return &maritalStatusServiceImpl{
		repo: repo,
	}
}

func (s *maritalStatusServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.MaritalStatusResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch marital status: %v", err)
		return nil, nil, err
	}

	var response []res.MaritalStatusResponse
	for _, item := range items {
		response = append(response, res.MaritalStatusResponse{
			ID:            item.UUID,
			MaritalStatus: item.MaterialStatus,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *maritalStatusServiceImpl) GetByID(ctx context.Context, id string) (*res.MaritalStatusResponse, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Marital status not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.MaritalStatusResponse{
		ID:            item.UUID,
		MaritalStatus: item.MaterialStatus,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}, nil
}

func (s *maritalStatusServiceImpl) Create(ctx context.Context, request req.CreateMaritalStatusRequest) (*res.MaritalStatusResponse, error) {
	item := &model.MaritalStatus{
		MaterialStatus: request.MaritalStatus,
	}

	err := s.repo.Create(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to create marital status: %v", err)
		return nil, err
	}

	return &res.MaritalStatusResponse{
		ID:            item.UUID,
		MaritalStatus: item.MaterialStatus,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}, nil
}

func (s *maritalStatusServiceImpl) Update(ctx context.Context, id string, request req.UpdateMaritalStatusRequest) (*res.MaritalStatusResponse, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Marital status not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	item.MaterialStatus = request.MaritalStatus

	err = s.repo.Update(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to update marital status: %v", err)
		return nil, err
	}

	return &res.MaritalStatusResponse{
		ID:            item.UUID,
		MaritalStatus: item.MaterialStatus,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}, nil
}

func (s *maritalStatusServiceImpl) Delete(ctx context.Context, id string) error {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Marital status not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, item.ID)
	if err != nil {
		logrus.Errorf("Failed to delete marital status: %v", err)
		return err
	}

	return nil
}
