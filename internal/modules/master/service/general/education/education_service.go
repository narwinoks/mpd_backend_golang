package education

import (
	req "backend-app/internal/modules/master/request/general/education"
	res "backend-app/internal/modules/master/response/general/education"
	"backend-app/pkg/pagination"
	"context"
)

type EducationService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.EducationResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.EducationResponse, error)
	Create(ctx context.Context, request req.CreateEducationRequest) (*res.EducationResponse, error)
	Update(ctx context.Context, id string, request req.UpdateEducationRequest) (*res.EducationResponse, error)
	Delete(ctx context.Context, id string) error
}
