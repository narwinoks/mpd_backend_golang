package village

import (
	"backend-app/internal/modules/master/model/location"
	"backend-app/pkg/pagination"
	"context"
)

type VillageRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]location.Village, int64, error)
	FindByID(ctx context.Context, id uint32) (*location.Village, error)
	Create(ctx context.Context, village *location.Village) error
	Update(ctx context.Context, village *location.Village) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*location.Village, error)
}
