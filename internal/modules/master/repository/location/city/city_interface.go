package city

import (
	"backend-app/internal/modules/master/model/location"
	"backend-app/pkg/pagination"
	"context"
)

type CityRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]location.City, int64, error)
	FindByID(ctx context.Context, id uint32) (*location.City, error)
	Create(ctx context.Context, city *location.City) error
	Update(ctx context.Context, city *location.City) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*location.City, error)
}
