package permission

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/permission"
	repo "backend-app/internal/modules/master/repository/permission"
	req "backend-app/internal/modules/master/request/permission"
	res "backend-app/internal/modules/master/response/permission"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type permissionServiceImpl struct {
	repo repo.PermissionRepository
}

func NewPermissionService(repo repo.PermissionRepository) PermissionService {
	return &permissionServiceImpl{
		repo: repo,
	}
}

func (s *permissionServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.PermissionResponse, *pagination.Meta, error) {
	permissions, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch permissions: %v", err)
		return nil, nil, err
	}

	var response []res.PermissionResponse
	for _, p := range permissions {
		response = append(response, res.PermissionResponse{
			ID:         p.UUID,
			Permission: p.Permission,
			CreatedAt:  p.CreatedAt,
			UpdatedAt:  p.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *permissionServiceImpl) GetByID(ctx context.Context, id string) (*res.PermissionResponse, error) {
	p, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Permission not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.PermissionResponse{
		ID:         p.UUID,
		Permission: p.Permission,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}, nil
}

func (s *permissionServiceImpl) Create(ctx context.Context, request req.CreatePermissionRequest) (*res.PermissionResponse, error) {
	p := &model.Permission{
		Permission: request.Permission,
	}

	err := s.repo.Create(ctx, p)
	if err != nil {
		logrus.Errorf("Failed to create permission: %v", err)
		return nil, err
	}

	return &res.PermissionResponse{
		ID:         p.UUID,
		Permission: p.Permission,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}, nil
}

func (s *permissionServiceImpl) Update(ctx context.Context, id string, request req.UpdatePermissionRequest) (*res.PermissionResponse, error) {
	p, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Permission not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	p.Permission = request.Permission

	err = s.repo.Update(ctx, p)
	if err != nil {
		logrus.Errorf("Failed to update permission: %v", err)
		return nil, err
	}

	return &res.PermissionResponse{
		ID:         p.UUID,
		Permission: p.Permission,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}, nil
}

func (s *permissionServiceImpl) Delete(ctx context.Context, id string) error {
	p, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Permission not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, p.ID)
	if err != nil {
		logrus.Errorf("Failed to delete permission: %v", err)
		return err
	}

	return nil
}
