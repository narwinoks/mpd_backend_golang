package religion

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"
)

type ReligionRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]general.Religion, int64, error)
	FindByID(ctx context.Context, id uint32) (*general.Religion, error)
	Create(ctx context.Context, religion *general.Religion) error
	Update(ctx context.Context, religion *general.Religion) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, uuid string) (*general.Religion, error)
}
