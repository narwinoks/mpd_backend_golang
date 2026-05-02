package bed

import (
	"backend-app/internal/modules/master/model/department"
	"backend-app/pkg/pagination"
	"context"
)

type BedRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]department.Bed, int64, error)
	FindByID(ctx context.Context, id uint32) (*department.Bed, error)
	FindByUuid(ctx context.Context, uuid string) (*department.Bed, error)
	Create(ctx context.Context, m *department.Bed) error
	Update(ctx context.Context, m *department.Bed) error
	Delete(ctx context.Context, id uint32) error
}
