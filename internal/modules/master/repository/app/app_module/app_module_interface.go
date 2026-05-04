package app_module

import (
	"backend-app/internal/modules/auth/models"
	"backend-app/pkg/pagination"
	"context"
)

type AppModuleRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]models.AppModule, int64, error)
	FindByID(ctx context.Context, id uint32) (*models.AppModule, error)
	Create(ctx context.Context, appModule *models.AppModule) error
	Update(ctx context.Context, appModule *models.AppModule) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, uuid string) (*models.AppModule, error)
}
