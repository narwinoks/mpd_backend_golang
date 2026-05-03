package employee

import (
	model "backend-app/internal/modules/master/model/employee"
	"backend-app/pkg/pagination"
	"context"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]model.Employee, int64, error)
	FindByID(ctx context.Context, id uint32) (*model.Employee, error)
	Create(ctx context.Context, employee *model.Employee) error
	Update(ctx context.Context, employee *model.Employee) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*model.Employee, error)
}
