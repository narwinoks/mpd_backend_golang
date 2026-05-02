package ward

import (
	req "backend-app/internal/modules/master/request/department/ward"
	res "backend-app/internal/modules/master/response/department/ward"
	"backend-app/pkg/pagination"
	"context"
)

type WardService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.WardResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.WardResponse, error)
	Create(ctx context.Context, request req.CreateWardRequest) (*res.WardResponse, error)
	Update(ctx context.Context, id string, request req.UpdateWardRequest) (*res.WardResponse, error)
	Delete(ctx context.Context, id string) error
}
