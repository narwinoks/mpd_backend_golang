package role

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/role"
	repo "backend-app/internal/modules/master/repository/role"
	req "backend-app/internal/modules/master/request/role"
	res "backend-app/internal/modules/master/response/role"
	"backend-app/pkg/pagination"

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

func (s *roleServiceImpl) GetAll(request pagination.Request) ([]res.RoleResponse, *pagination.Meta, error) {
	roles, total, err := s.repo.FindAll(request)
	if err != nil {
		logrus.Errorf("Failed to fetch roles: %v", err)
		return nil, nil, err
	}

	var response []res.RoleResponse
	for _, r := range roles {
		response = append(response, res.RoleResponse{
			ID:        r.ID,
			Role:      r.Role,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *roleServiceImpl) GetByID(id uint32) (*res.RoleResponse, error) {
	roleEntity, err := s.repo.FindByID(id)
	if err != nil {
		logrus.Errorf("Role not found with id %d: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.RoleResponse{
		ID:        roleEntity.ID,
		Role:      roleEntity.Role,
		CreatedAt: roleEntity.CreatedAt,
		UpdatedAt: roleEntity.UpdatedAt,
	}, nil
}

func (s *roleServiceImpl) Create(request req.CreateRoleRequest) (*res.RoleResponse, error) {
	roleEntity := &model.Role{
		Role: request.Role,
	}

	err := s.repo.Create(roleEntity)
	if err != nil {
		logrus.Errorf("Failed to create role: %v", err)
		return nil, err
	}

	return &res.RoleResponse{
		ID:        roleEntity.ID,
		Role:      roleEntity.Role,
		CreatedAt: roleEntity.CreatedAt,
		UpdatedAt: roleEntity.UpdatedAt,
	}, nil
}

func (s *roleServiceImpl) Update(id uint32, request req.UpdateRoleRequest) (*res.RoleResponse, error) {
	roleEntity, err := s.repo.FindByID(id)
	if err != nil {
		logrus.Errorf("Role not found for update with id %d: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	roleEntity.Role = request.Role

	err = s.repo.Update(roleEntity)
	if err != nil {
		logrus.Errorf("Failed to update role: %v", err)
		return nil, err
	}

	return &res.RoleResponse{
		ID:        roleEntity.ID,
		Role:      roleEntity.Role,
		CreatedAt: roleEntity.CreatedAt,
		UpdatedAt: roleEntity.UpdatedAt,
	}, nil
}

func (s *roleServiceImpl) Delete(id uint32) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		logrus.Errorf("Role not found for deletion with id %d: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}

	err = s.repo.Delete(id)
	if err != nil {
		logrus.Errorf("Failed to delete role: %v", err)
		return err
	}

	return nil
}
