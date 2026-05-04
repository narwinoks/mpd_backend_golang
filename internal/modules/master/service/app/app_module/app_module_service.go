package app_module

import (
	req "backend-app/internal/modules/master/request/app/app_module"
	res "backend-app/internal/modules/master/response/app/app_module"
	"backend-app/pkg/pagination"
	"context"
)

type AppModuleService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.AppModuleResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.AppModuleResponse, error)
	Create(ctx context.Context, request req.CreateAppModuleRequest) (string, error)
	Update(ctx context.Context, id string, request req.UpdateAppModuleRequest) (string, error)
	Delete(ctx context.Context, id string) error
}
