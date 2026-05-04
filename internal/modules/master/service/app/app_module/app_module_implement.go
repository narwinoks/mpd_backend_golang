package app_module

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/models"
	repo "backend-app/internal/modules/master/repository/app/app_module"
	req "backend-app/internal/modules/master/request/app/app_module"
	res "backend-app/internal/modules/master/response/app/app_module"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type appModuleServiceImpl struct {
	repo repo.AppModuleRepository
}

func NewAppModuleService(repo repo.AppModuleRepository) AppModuleService {
	return &appModuleServiceImpl{
		repo: repo,
	}
}

func (s *appModuleServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.AppModuleResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch app modules: %v", err)
		return nil, nil, err
	}

	var response []res.AppModuleResponse
	for _, item := range items {
		response = append(response, res.AppModuleResponse{
			ID:        item.UUID,
			Code:      item.Code,
			Name:      item.Name,
			Category:  item.Category,
			SortOrder: item.SortOrder,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *appModuleServiceImpl) GetByID(ctx context.Context, id string) (*res.AppModuleResponse, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("App module not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.AppModuleResponse{
		ID:        item.UUID,
		Code:      item.Code,
		Name:      item.Name,
		Category:  item.Category,
		SortOrder: item.SortOrder,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}, nil
}

func (s *appModuleServiceImpl) Create(ctx context.Context, request req.CreateAppModuleRequest) (string, error) {
	item := &models.AppModule{
		Code:      request.Code,
		Name:      request.Name,
		Category:  request.Category,
		SortOrder: request.SortOrder,
	}

	err := s.repo.Create(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to create app module: %v", err)
		return "", err
	}

	return item.UUID, nil
}

func (s *appModuleServiceImpl) Update(ctx context.Context, id string, request req.UpdateAppModuleRequest) (string, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("App module not found for update with id %s: %v", id, err)
		return "", exception.NewNotFoundError("Data not found")
	}

	item.Code = request.Code
	item.Name = request.Name
	item.Category = request.Category
	item.SortOrder = request.SortOrder

	err = s.repo.Update(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to update app module: %v", err)
		return "", err
	}

	return item.UUID, nil
}

func (s *appModuleServiceImpl) Delete(ctx context.Context, id string) error {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("App module not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, item.ID)
	if err != nil {
		logrus.Errorf("Failed to delete app module: %v", err)
		return err
	}

	return nil
}
