package registry

import (
	req "backend-app/internal/modules/master/request/registry"
	res "backend-app/internal/modules/master/response/registry"
	"backend-app/pkg/pagination"
	"context"
)

type RegistryService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.RegistryResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.RegistryResponse, error)
	Create(ctx context.Context, request req.CreateRegistryRequest) (*res.RegistryResponse, error)
	Update(ctx context.Context, id string, request req.UpdateRegistryRequest) (*res.RegistryResponse, error)
	Delete(ctx context.Context, id string) error
	GetNestedMenu(ctx context.Context) ([]res.RegistryResponse, error)
}
