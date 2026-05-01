package gender

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"
)

type GenderRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]general.Gender, int64, error)
	FindByID(ctx context.Context, id uint32) (*general.Gender, error)
	Create(ctx context.Context, gender *general.Gender) error
	Update(ctx context.Context, gender *general.Gender) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, uuid string) (*general.Gender, error)
}
