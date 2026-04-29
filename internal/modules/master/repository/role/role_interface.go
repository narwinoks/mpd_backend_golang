package role

import (
	"backend-app/internal/modules/master/model/role"
	"backend-app/pkg/pagination"
	"context"
)

type RoleRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]role.Role, int64, error)
	FindByID(ctx context.Context, id uint32) (*role.Role, error)
	Create(ctx context.Context, role *role.Role) error
	Update(ctx context.Context, role *role.Role) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*role.Role, error)
}
