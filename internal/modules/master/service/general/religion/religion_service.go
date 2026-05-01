package religion

import (
	req "backend-app/internal/modules/master/request/general/religion"
	res "backend-app/internal/modules/master/response/general/religion"
	"backend-app/pkg/pagination"
	"context"
)

type ReligionService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.ReligionResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.ReligionResponse, error)
	Create(ctx context.Context, request req.CreateReligionRequest) (*res.ReligionResponse, error)
	Update(ctx context.Context, id string, request req.UpdateReligionRequest) (*res.ReligionResponse, error)
	Delete(ctx context.Context, id string) error
}
