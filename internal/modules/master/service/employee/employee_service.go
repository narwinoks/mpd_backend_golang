package employee

import (
	req "backend-app/internal/modules/master/request/employee"
	res "backend-app/internal/modules/master/response/employee"
	"backend-app/pkg/pagination"
	"context"
)

type EmployeeService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.EmployeeListResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.EmployeeResponse, error)
	Create(ctx context.Context, request req.CreateEmployeeRequest) (string, error)
	Update(ctx context.Context, id string, request req.UpdateEmployeeRequest) (string, error)
	Delete(ctx context.Context, id string) error
}
