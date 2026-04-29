package permission

import (
	"backend-app/internal/modules/master/model/permission"
	"backend-app/pkg/pagination"
	"context"
)

type PermissionRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]permission.Permission, int64, error)
	FindByID(ctx context.Context, id uint32) (*permission.Permission, error)
	Create(ctx context.Context, p *permission.Permission) error
	Update(ctx context.Context, p *permission.Permission) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, uuid string) (*permission.Permission, error)
}
