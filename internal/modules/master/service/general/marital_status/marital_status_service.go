package marital_status

import (
	req "backend-app/internal/modules/master/request/general/marital_status"
	res "backend-app/internal/modules/master/response/general/marital_status"
	"backend-app/pkg/pagination"
	"context"
)

type MaritalStatusService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.MaritalStatusResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.MaritalStatusResponse, error)
	Create(ctx context.Context, request req.CreateMaritalStatusRequest) (*res.MaritalStatusResponse, error)
	Update(ctx context.Context, id string, request req.UpdateMaritalStatusRequest) (*res.MaritalStatusResponse, error)
	Delete(ctx context.Context, id string) error
}
