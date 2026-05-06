package app_menu

import (
	req "backend-app/internal/modules/master/request/app_menu"
	res "backend-app/internal/modules/master/response/app_menu"
	"backend-app/pkg/pagination"
	"context"
)

type AppMenuService interface {
	GetAll(ctx context.Context, request req.AppMenuFilterRequest) ([]res.AppMenuResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.AppMenuResponse, error)
	Create(ctx context.Context, request req.CreateAppMenuRequest) (string, error)
	Update(ctx context.Context, id string, request req.UpdateAppMenuRequest) (string, error)
	Delete(ctx context.Context, id string) error
}
