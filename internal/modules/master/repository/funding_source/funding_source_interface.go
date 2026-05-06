package funding_source

import (
	model "backend-app/internal/modules/master/model/funding_source"
	"backend-app/pkg/pagination"
	"context"
)

type FundingSourceRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]model.FundingSource, int64, error)
	FindByID(ctx context.Context, id uint32) (*model.FundingSource, error)
	FindByUUID(ctx context.Context, uuid string) (*model.FundingSource, error)
	Create(ctx context.Context, fundingSource *model.FundingSource) error
	Update(ctx context.Context, fundingSource *model.FundingSource) error
	Delete(ctx context.Context, id uint32) error
}
