package app_menu

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/models"
	repoAppModule "backend-app/internal/modules/master/repository/app/app_module"
	repo "backend-app/internal/modules/master/repository/app_menu"
	req "backend-app/internal/modules/master/request/app_menu"
	res "backend-app/internal/modules/master/response/app_menu"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type appMenuServiceImpl struct {
	repo          repo.AppMenuRepository
	appModuleRepo repoAppModule.AppModuleRepository
}

func NewAppMenuService(repo repo.AppMenuRepository, appModuleRepo repoAppModule.AppModuleRepository) AppMenuService {
	return &appMenuServiceImpl{
		repo:          repo,
		appModuleRepo: appModuleRepo,
	}
}

func (s *appMenuServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.AppMenuResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch app menus: %v", err)
		return nil, nil, err
	}

	var response []res.AppMenuResponse
	for _, item := range items {
		var parentUUID *string
		var parentName *string
		if item.Parent != nil {
			parentUUID = &item.Parent.UUID
			parentName = &item.Parent.Name
		}

		response = append(response, res.AppMenuResponse{
			ID:            item.UUID,
			AppModuleID:   item.AppModule.UUID,
			AppModuleName: item.AppModule.Name,
			ParentID:      parentUUID,
			ParentName:    parentName,
			Code:          item.Code,
			Name:          item.Name,
			Path:          item.Path,
			Description:   item.Description,
			Icon:          item.Icon,
			SortOrder:     item.SortOrder,
			IsActive:      item.IsActive,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *appMenuServiceImpl) GetByID(ctx context.Context, id string) (*res.AppMenuResponse, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("App menu not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	var parentUUID *string
	var parentName *string
	if item.Parent != nil {
		parentUUID = &item.Parent.UUID
		parentName = &item.Parent.Name
	}

	return &res.AppMenuResponse{
		ID:            item.UUID,
		AppModuleID:   item.AppModule.UUID,
		AppModuleName: item.AppModule.Name,
		ParentID:      parentUUID,
		ParentName:    parentName,
		Code:          item.Code,
		Name:          item.Name,
		Path:          item.Path,
		Description:   item.Description,
		Icon:          item.Icon,
		SortOrder:     item.SortOrder,
		IsActive:      item.IsActive,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}, nil
}

func (s *appMenuServiceImpl) Create(ctx context.Context, request req.CreateAppMenuRequest) (string, error) {
	appModule, err := s.appModuleRepo.FindByUuid(ctx, request.AppModuleID)
	if err != nil {
		return "", exception.NewNotFoundError("App Module not found")
	}

	var parentID *uint32
	if request.ParentID != nil && *request.ParentID != "" {
		parent, err := s.repo.FindByUuid(ctx, *request.ParentID)
		if err != nil {
			return "", exception.NewNotFoundError("Parent Menu not found")
		}
		parentID = &parent.ID
	}

	item := &models.AppMenu{
		AppModuleID: appModule.ID,
		ParentID:    parentID,
		Code:        request.Code,
		Name:        request.Name,
		Path:        request.Path,
		Description: request.Description,
		Icon:        request.Icon,
		SortOrder:   request.SortOrder,
	}

	err = s.repo.Create(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to create app menu: %v", err)
		return "", err
	}

	return item.UUID, nil
}

func (s *appMenuServiceImpl) Update(ctx context.Context, id string, request req.UpdateAppMenuRequest) (string, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("App menu not found for update with id %s: %v", id, err)
		return "", exception.NewNotFoundError("Data not found")
	}

	appModule, err := s.appModuleRepo.FindByUuid(ctx, request.AppModuleID)
	if err != nil {
		return "", exception.NewNotFoundError("App Module not found")
	}

	var parentID *uint32
	if request.ParentID != nil && *request.ParentID != "" {
		parent, err := s.repo.FindByUuid(ctx, *request.ParentID)
		if err != nil {
			return "", exception.NewNotFoundError("Parent Menu not found")
		}
		parentID = &parent.ID
	}

	item.AppModuleID = appModule.ID
	item.ParentID = parentID
	item.Code = request.Code
	item.Name = request.Name
	item.Path = request.Path
	item.Description = request.Description
	item.Icon = request.Icon
	item.SortOrder = request.SortOrder

	if request.IsActive != nil {
		item.IsActive = *request.IsActive
	}

	err = s.repo.Update(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to update app menu: %v", err)
		return "", err
	}

	return item.UUID, nil
}

func (s *appMenuServiceImpl) Delete(ctx context.Context, id string) error {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("App menu not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, item.ID)
	if err != nil {
		logrus.Errorf("Failed to delete app menu: %v", err)
		return err
	}

	return nil
}
