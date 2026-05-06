package funding_source

import (
	req "backend-app/internal/modules/master/request/funding_source"
	res "backend-app/internal/modules/master/response/funding_source"
	"backend-app/pkg/pagination"
	"context"
)

type FundingSourceService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.FundingSourceResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.FundingSourceResponse, error)
	Create(ctx context.Context, request req.CreateFundingSourceRequest) (*res.FundingSourceResponse, error)
	Update(ctx context.Context, id string, request req.UpdateFundingSourceRequest) (*res.FundingSourceResponse, error)
	Delete(ctx context.Context, id string) error
}
