package registry

import (
	"backend-app/internal/modules/master/model/registry"
	"backend-app/pkg/pagination"
	"context"
)

type RegistryRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]registry.Registry, int64, error)
	FindByID(ctx context.Context, id uint32) (*registry.Registry, error)
	Create(ctx context.Context, reg *registry.Registry) error
	Update(ctx context.Context, reg *registry.Registry) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, uuid string) (*registry.Registry, error)
	FindNested(ctx context.Context) ([]registry.Registry, error)
}
