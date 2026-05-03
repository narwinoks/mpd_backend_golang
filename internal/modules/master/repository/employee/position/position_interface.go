package position

import (
	"backend-app/internal/modules/master/model/job"
	"backend-app/pkg/pagination"
	"context"
)

type PositionRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]job.Position, int64, error)
	FindByID(ctx context.Context, id uint32) (*job.Position, error)
	Create(ctx context.Context, position *job.Position) error
	Update(ctx context.Context, position *job.Position) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*job.Position, error)
}
