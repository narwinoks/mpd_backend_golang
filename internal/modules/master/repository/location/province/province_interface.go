package province

import (
	"backend-app/internal/modules/master/model/location"
	"backend-app/pkg/pagination"
	"context"
)

type ProvinceRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]location.Province, int64, error)
	FindByID(ctx context.Context, id uint32) (*location.Province, error)
	Create(ctx context.Context, province *location.Province) error
	Update(ctx context.Context, province *location.Province) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*location.Province, error)
}
