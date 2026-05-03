package position

import (
	req "backend-app/internal/modules/master/request/employee/position"
	res "backend-app/internal/modules/master/response/employee/position"
	"backend-app/pkg/pagination"
	"context"
)

type PositionService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.PositionResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.PositionResponse, error)
	Create(ctx context.Context, request req.CreatePositionRequest) (*res.PositionResponse, error)
	Update(ctx context.Context, id string, request req.UpdatePositionRequest) (*res.PositionResponse, error)
	Delete(ctx context.Context, id string) error
}
