package subdistrict

import (
	req "backend-app/internal/modules/master/request/location/subdistrict"
	res "backend-app/internal/modules/master/response/location/subdistrict"
	"backend-app/pkg/pagination"
	"context"
)

type SubdistrictService interface {
	GetAll(ctx context.Context, request req.FindAllRequest) ([]res.SubdistrictResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.SubdistrictResponse, error)
	Create(ctx context.Context, request req.CreateSubdistrictRequest) (string, error)
	Update(ctx context.Context, id string, request req.UpdateSubdistrictRequest) (string, error)
	Delete(ctx context.Context, id string) error
}
