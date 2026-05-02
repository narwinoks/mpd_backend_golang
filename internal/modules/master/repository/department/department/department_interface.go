package department

import (
	"backend-app/internal/modules/master/model/department"
	"backend-app/pkg/pagination"
	"context"
)

type DepartmentRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]department.Department, int64, error)
	FindByID(ctx context.Context, id uint32) (*department.Department, error)
	FindByUuid(ctx context.Context, uuid string) (*department.Department, error)
	Create(ctx context.Context, m *department.Department) error
	Update(ctx context.Context, m *department.Department) error
	Delete(ctx context.Context, id uint32) error
}
