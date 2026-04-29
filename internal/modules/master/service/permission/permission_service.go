package permission

import (
	req "backend-app/internal/modules/master/request/permission"
	res "backend-app/internal/modules/master/response/permission"
	"backend-app/pkg/pagination"
	"context"
)

type PermissionService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.PermissionResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.PermissionResponse, error)
	Create(ctx context.Context, request req.CreatePermissionRequest) (*res.PermissionResponse, error)
	Update(ctx context.Context, id string, request req.UpdatePermissionRequest) (*res.PermissionResponse, error)
	Delete(ctx context.Context, id string) error
}
