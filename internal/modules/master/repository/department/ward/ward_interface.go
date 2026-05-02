package ward

import (
	"backend-app/internal/modules/master/model/department"
	"backend-app/pkg/pagination"
	"context"
)

type WardRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]department.Ward, int64, error)
	FindByID(ctx context.Context, id uint32) (*department.Ward, error)
	FindByUuid(ctx context.Context, uuid string) (*department.Ward, error)
	Create(ctx context.Context, m *department.Ward) error
	Update(ctx context.Context, m *department.Ward) error
	Delete(ctx context.Context, id uint32) error
}
