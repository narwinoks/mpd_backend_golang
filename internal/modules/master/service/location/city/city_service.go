package city

import (
	req "backend-app/internal/modules/master/request/location/city"
	res "backend-app/internal/modules/master/response/location/city"
	"backend-app/pkg/pagination"
	"context"
)

type CityService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.CityResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.CityResponse, error)
	Create(ctx context.Context, request req.CreateCityRequest) (string, error)
	Update(ctx context.Context, id string, request req.UpdateCityRequest) (string, error)
	Delete(ctx context.Context, id string) error
}
