package village

import (
	req "backend-app/internal/modules/master/request/location/village"
	res "backend-app/internal/modules/master/response/location/village"
	"backend-app/pkg/pagination"
	"context"
)

type VillageService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.VillageResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.VillageResponse, error)
	Create(ctx context.Context, request req.CreateVillageRequest) (string, error)
	Update(ctx context.Context, id string, request req.UpdateVillageRequest) (string, error)
	Delete(ctx context.Context, id string) error
}
