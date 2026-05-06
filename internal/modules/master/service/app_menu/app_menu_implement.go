package app_menu

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/models"
	repoAppModule "backend-app/internal/modules/master/repository/app/app_module"
	repo "backend-app/internal/modules/master/repository/app_menu"
	req "backend-app/internal/modules/master/request/app_menu"
	resAppModule "backend-app/internal/modules/master/response/app/app_module"
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

func (s *appMenuServiceImpl) GetAll(ctx context.Context, request req.AppMenuFilterRequest) ([]res.AppMenuResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch app menus: %v", err)
		return nil, nil, err
	}

	var response []res.AppMenuResponse
	for _, item := range items {
		response = append(response, s.mapToResponse(item))
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

	response := s.mapToResponse(*item)
	return &response, nil
}

func (s *appMenuServiceImpl) mapToResponse(item models.AppMenu) res.AppMenuResponse {
	var parentUUID *string
	var parentName *string
	var parentObj *res.AppMenuResponse

	if item.Parent != nil {
		parentUUID = &item.Parent.UUID
		parentName = &item.Parent.Name
		parentObj = &res.AppMenuResponse{
			ID:          item.Parent.ID,
			UUID:        item.Parent.UUID,
			Code:        item.Parent.Code,
			Name:        item.Parent.Name,
			Path:        item.Parent.Path,
			Description: item.Parent.Description,
			Icon:        item.Parent.Icon,
			SortOrder:   item.Parent.SortOrder,
			IsActive:    item.Parent.IsActive,
			CreatedAt:   item.Parent.CreatedAt,
			UpdatedAt:   item.Parent.UpdatedAt,
		}
	}

	var appModuleObj *resAppModule.AppModuleResponse
	if item.AppModule.ID != 0 {
		appModuleObj = &resAppModule.AppModuleResponse{
			ID:        item.AppModule.UUID,
			Code:      item.AppModule.Code,
			Name:      item.AppModule.Name,
			Category:  item.AppModule.Category,
			SortOrder: item.AppModule.SortOrder,
			CreatedAt: item.AppModule.CreatedAt,
			UpdatedAt: item.AppModule.UpdatedAt,
		}
	}

	response := res.AppMenuResponse{
		ID:            item.ID,
		UUID:          item.UUID,
		AppModuleID:   item.AppModule.UUID,
		AppModuleName: item.AppModule.Name,
		AppModule:     appModuleObj,
		ParentID:      parentUUID,
		ParentName:    parentName,
		Parent:        parentObj,
		Code:          item.Code,
		Name:          item.Name,
		Path:          item.Path,
		Description:   item.Description,
		Icon:          item.Icon,
		SortOrder:     item.SortOrder,
		IsActive:      item.IsActive,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}

	if len(item.SubMenus) > 0 {
		response.Children = make([]*res.AppMenuResponse, 0)
		for _, sub := range item.SubMenus {
			mappedSub := s.mapToResponse(*sub)
			response.Children = append(response.Children, &mappedSub)
		}
	}

	return response
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
