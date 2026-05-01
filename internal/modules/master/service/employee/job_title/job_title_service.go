package job_title

import (
	req "backend-app/internal/modules/master/request/employee/job_title"
	res "backend-app/internal/modules/master/response/employee/job_title"
	"backend-app/pkg/pagination"
	"context"
)

type JobTitleService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.JobTitleResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.JobTitleResponse, error)
	Create(ctx context.Context, request req.CreateJobTitleRequest) (*res.JobTitleResponse, error)
	Update(ctx context.Context, id string, request req.UpdateJobTitleRequest) (*res.JobTitleResponse, error)
	Delete(ctx context.Context, id string) error
}
