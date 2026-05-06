package subdistrict

import (
	"backend-app/internal/modules/master/model/location"
	req "backend-app/internal/modules/master/request/location/subdistrict"
	"context"
)

type SubdistrictRepository interface {
	FindAll(ctx context.Context, req req.FindAllRequest) ([]location.Subdistrict, int64, error)
	FindByID(ctx context.Context, id uint32) (*location.Subdistrict, error)
	Create(ctx context.Context, subdistrict *location.Subdistrict) error
	Update(ctx context.Context, subdistrict *location.Subdistrict) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*location.Subdistrict, error)
}
