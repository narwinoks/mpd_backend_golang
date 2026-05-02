package department

import (
	req "backend-app/internal/modules/master/request/department/department"
	res "backend-app/internal/modules/master/response/department/department"
	"backend-app/pkg/pagination"
	"context"
)

type DepartmentService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.DepartmentResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.DepartmentResponse, error)
	Create(ctx context.Context, request req.CreateDepartmentRequest) (*res.DepartmentResponse, error)
	Update(ctx context.Context, id string, request req.UpdateDepartmentRequest) (*res.DepartmentResponse, error)
	Delete(ctx context.Context, id string) error
}
