package employment_status

import (
	req "backend-app/internal/modules/master/request/employee/employment_status"
	res "backend-app/internal/modules/master/response/employee/employment_status"
	"backend-app/pkg/pagination"
	"context"
)

type EmploymentStatusService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.EmploymentStatusResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.EmploymentStatusResponse, error)
	Create(ctx context.Context, request req.CreateEmploymentStatusRequest) (*res.EmploymentStatusResponse, error)
	Update(ctx context.Context, id string, request req.UpdateEmploymentStatusRequest) (*res.EmploymentStatusResponse, error)
	Delete(ctx context.Context, id string) error
}
