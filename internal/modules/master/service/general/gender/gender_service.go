package gender

import (
	req "backend-app/internal/modules/master/request/general/gender"
	res "backend-app/internal/modules/master/response/general/gender"
	"backend-app/pkg/pagination"
	"context"
)

type GenderService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.GenderResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.GenderResponse, error)
	Create(ctx context.Context, request req.CreateGenderRequest) (*res.GenderResponse, error)
	Update(ctx context.Context, id string, request req.UpdateGenderRequest) (*res.GenderResponse, error)
	Delete(ctx context.Context, id string) error
}
