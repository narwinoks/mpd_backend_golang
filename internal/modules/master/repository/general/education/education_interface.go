package education

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"
)

type EducationRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]general.Education, int64, error)
	FindByID(ctx context.Context, id uint32) (*general.Education, error)
	Create(ctx context.Context, education *general.Education) error
	Update(ctx context.Context, education *general.Education) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*general.Education, error)
}
