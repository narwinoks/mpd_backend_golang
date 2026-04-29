package role

import (
	req "backend-app/internal/modules/master/request/role"
	res "backend-app/internal/modules/master/response/role"
	"backend-app/pkg/pagination"
	"context"
)

type RoleService interface {
	GetAll(ctx context.Context, request pagination.Request) ([]res.RoleResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.RoleResponse, error)
	Create(ctx context.Context, request req.CreateRoleRequest) (*res.RoleResponse, error)
	Update(ctx context.Context, id string, request req.UpdateRoleRequest) (*res.RoleResponse, error)
	Delete(ctx context.Context, id string) error
}
