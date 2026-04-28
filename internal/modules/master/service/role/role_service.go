package role

import (
	req "backend-app/internal/modules/master/request/role"
	res "backend-app/internal/modules/master/response/role"
	"backend-app/pkg/pagination"
)

type RoleService interface {
	GetAll(request pagination.Request) ([]res.RoleResponse, *pagination.Meta, error)
	GetByID(id uint32) (*res.RoleResponse, error)
	Create(request req.CreateRoleRequest) (*res.RoleResponse, error)
	Update(id uint32, request req.UpdateRoleRequest) (*res.RoleResponse, error)
	Delete(id uint32) error
}
