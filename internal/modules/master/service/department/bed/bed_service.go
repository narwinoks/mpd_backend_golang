package bed

import (
	req "backend-app/internal/modules/master/request/department/bed"
	res "backend-app/internal/modules/master/response/department/bed"
	"backend-app/pkg/pagination"
	"context"
)

type BedService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.BedResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.BedResponse, error)
	Create(ctx context.Context, request req.CreateBedRequest) (*res.BedResponse, error)
	Update(ctx context.Context, id string, request req.UpdateBedRequest) (*res.BedResponse, error)
	Delete(ctx context.Context, id string) error
}
