package app_menu

import (
	"backend-app/internal/modules/auth/models"
	"backend-app/pkg/pagination"
	"context"
)

type AppMenuRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]models.AppMenu, int64, error)
	FindByID(ctx context.Context, id uint32) (*models.AppMenu, error)
	FindByUuid(ctx context.Context, uuid string) (*models.AppMenu, error)
	Create(ctx context.Context, item *models.AppMenu) error
	Update(ctx context.Context, item *models.AppMenu) error
	Delete(ctx context.Context, id uint32) error
}
