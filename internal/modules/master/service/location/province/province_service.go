package province

import (
	req "backend-app/internal/modules/master/request/location/province"
	res "backend-app/internal/modules/master/response/location/province"
	"backend-app/pkg/pagination"
	"context"
)

type ProvinceService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.ProvinceResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.ProvinceResponse, error)
	Create(ctx context.Context, request req.CreateProvinceRequest) (*res.ProvinceResponse, error)
	Update(ctx context.Context, id string, request req.UpdateProvinceRequest) (*res.ProvinceResponse, error)
	Delete(ctx context.Context, id string) error
}
