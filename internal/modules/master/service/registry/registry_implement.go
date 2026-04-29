package registry

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/registry"
	repo "backend-app/internal/modules/master/repository/registry"
	req "backend-app/internal/modules/master/request/registry"
	res "backend-app/internal/modules/master/response/registry"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type registryServiceImpl struct {
	repo repo.RegistryRepository
}

func NewRegistryService(repo repo.RegistryRepository) RegistryService {
	return &registryServiceImpl{
		repo: repo,
	}
}

func (s *registryServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.RegistryResponse, *pagination.Meta, error) {
	registries, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch registries: %v", err)
		return nil, nil, err
	}

	var response []res.RegistryResponse
	for _, r := range registries {
		response = append(response, res.RegistryResponse{
			ID:        r.UUID,
			Name:      r.Name,
			Path:      r.Path,
			Icon:      r.Icon,
			HeadID:    r.HeadID,
			SortOrder: r.SortOrder,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *registryServiceImpl) GetByID(ctx context.Context, id string) (*res.RegistryResponse, error) {
	reg, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Registry not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.RegistryResponse{
		ID:        reg.UUID,
		Name:      reg.Name,
		Path:      reg.Path,
		Icon:      reg.Icon,
		HeadID:    reg.HeadID,
		SortOrder: reg.SortOrder,
		CreatedAt: reg.CreatedAt,
		UpdatedAt: reg.UpdatedAt,
	}, nil
}

func (s *registryServiceImpl) Create(ctx context.Context, request req.CreateRegistryRequest) (*res.RegistryResponse, error) {
	var headIDPtr *uint32
	if request.HeadID != "" {
		uuid, err := s.repo.FindByUuid(ctx, request.HeadID)
		if err == nil && uuid != nil {
			headIDPtr = &uuid.ID
		}
	}
	reg := &model.Registry{
		Name:      request.Name,
		Path:      request.Path,
		Icon:      request.Icon,
		HeadID:    headIDPtr,
		SortOrder: request.SortOrder,
	}

	err := s.repo.Create(ctx, reg)
	if err != nil {
		logrus.Errorf("Failed to create registry: %v", err)
		return nil, err
	}

	return &res.RegistryResponse{
		ID:        reg.UUID,
		Name:      reg.Name,
		Path:      reg.Path,
		Icon:      reg.Icon,
		HeadID:    reg.HeadID,
		SortOrder: reg.SortOrder,
		CreatedAt: reg.CreatedAt,
		UpdatedAt: reg.UpdatedAt,
	}, nil
}

func (s *registryServiceImpl) Update(ctx context.Context, id string, request req.UpdateRegistryRequest) (*res.RegistryResponse, error) {
	reg, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Registry not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}
	var headIDPtr *uint32
	if request.HeadID != "" {
		uuid, err := s.repo.FindByUuid(ctx, request.HeadID)
		if err == nil && uuid != nil {
			headIDPtr = &uuid.ID
		}
	}

	reg.Name = request.Name
	reg.Path = request.Path
	reg.Icon = request.Icon
	reg.HeadID = headIDPtr
	reg.SortOrder = request.SortOrder

	err = s.repo.Update(ctx, reg)
	if err != nil {
		logrus.Errorf("Failed to update registry: %v", err)
		return nil, err
	}

	return &res.RegistryResponse{
		ID:        reg.UUID,
		Name:      reg.Name,
		Path:      reg.Path,
		Icon:      reg.Icon,
		HeadID:    reg.HeadID,
		SortOrder: reg.SortOrder,
		CreatedAt: reg.CreatedAt,
		UpdatedAt: reg.UpdatedAt,
	}, nil
}

func (s *registryServiceImpl) Delete(ctx context.Context, id string) error {
	reg, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Registry not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, reg.ID)
	if err != nil {
		logrus.Errorf("Failed to delete registry: %v", err)
		return err
	}

	return nil
}

func (s *registryServiceImpl) GetNestedMenu(ctx context.Context) ([]res.RegistryResponse, error) {
	registries, err := s.repo.FindNested(ctx)
	if err != nil {
		logrus.Errorf("Failed to fetch nested registries: %v", err)
		return nil, err
	}

	var response []res.RegistryResponse
	for _, r := range registries {
		response = append(response, s.mapToResponse(r))
	}

	return response, nil
}

func (s *registryServiceImpl) mapToResponse(reg model.Registry) res.RegistryResponse {
	response := res.RegistryResponse{
		ID:        reg.UUID,
		Name:      reg.Name,
		Path:      reg.Path,
		Icon:      reg.Icon,
		HeadID:    reg.HeadID,
		SortOrder: reg.SortOrder,
		CreatedAt: reg.CreatedAt,
		UpdatedAt: reg.UpdatedAt,
	}

	if len(reg.Children) > 0 {
		for _, child := range reg.Children {
			response.Children = append(response.Children, s.mapToResponse(child))
		}
	}

	return response
}
