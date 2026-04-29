package role

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/role"
	repo "backend-app/internal/modules/master/repository/role"
	req "backend-app/internal/modules/master/request/role"
	res "backend-app/internal/modules/master/response/role"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type roleServiceImpl struct {
	repo repo.RoleRepository
}

func NewRoleService(repo repo.RoleRepository) RoleService {
	return &roleServiceImpl{
		repo: repo,
	}
}

func (s *roleServiceImpl) GetAll(ctx context.Context, request pagination.Request) ([]res.RoleResponse, *pagination.Meta, error) {
	roles, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch roles: %v", err)
		return nil, nil, err
	}

	var response []res.RoleResponse
	for _, r := range roles {
		response = append(response, res.RoleResponse{
			ID:        r.UUID,
			Role:      r.Role,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *roleServiceImpl) GetByID(ctx context.Context, id string) (*res.RoleResponse, error) {
	roleEntity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Role not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.RoleResponse{
		ID:        roleEntity.UUID,
		Role:      roleEntity.Role,
		CreatedAt: roleEntity.CreatedAt,
		UpdatedAt: roleEntity.UpdatedAt,
	}, nil
}

func (s *roleServiceImpl) Create(ctx context.Context, request req.CreateRoleRequest) (*res.RoleResponse, error) {
	roleEntity := &model.Role{
		Role: request.Role,
	}

	err := s.repo.Create(ctx, roleEntity)
	if err != nil {
		logrus.Errorf("Failed to create role: %v", err)
		return nil, err
	}

	return &res.RoleResponse{
		ID:        roleEntity.UUID,
		Role:      roleEntity.Role,
		CreatedAt: roleEntity.CreatedAt,
		UpdatedAt: roleEntity.UpdatedAt,
	}, nil
}

func (s *roleServiceImpl) Update(ctx context.Context, id string, request req.UpdateRoleRequest) (*res.RoleResponse, error) {
	roleEntity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Role not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	roleEntity.Role = request.Role

	err = s.repo.Update(ctx, roleEntity)
	if err != nil {
		logrus.Errorf("Failed to update role: %v", err)
		return nil, err
	}

	return &res.RoleResponse{
		ID:        roleEntity.UUID,
		Role:      roleEntity.Role,
		CreatedAt: roleEntity.CreatedAt,
		UpdatedAt: roleEntity.UpdatedAt,
	}, nil
}

func (s *roleServiceImpl) Delete(ctx context.Context, id string) error {
	role, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Role not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, role.ID)
	if err != nil {
		logrus.Errorf("Failed to delete role: %v", err)
		return err
	}

	return nil
}
