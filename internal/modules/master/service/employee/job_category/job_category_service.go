package job_category

import (
	req "backend-app/internal/modules/master/request/employee/job_category"
	res "backend-app/internal/modules/master/response/employee/job_category"
	"backend-app/pkg/pagination"
	"context"
)

type JobCategoryService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.JobCategoryResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.JobCategoryResponse, error)
	Create(ctx context.Context, request req.CreateJobCategoryRequest) (*res.JobCategoryResponse, error)
	Update(ctx context.Context, id string, request req.UpdateJobCategoryRequest) (*res.JobCategoryResponse, error)
	Delete(ctx context.Context, id string) error
}
