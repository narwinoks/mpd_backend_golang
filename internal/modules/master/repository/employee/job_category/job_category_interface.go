package job_category

import (
	"backend-app/internal/modules/master/model/job"
	"backend-app/pkg/pagination"
	"context"
)

type JobCategoryRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]job.JobCategory, int64, error)
	FindByID(ctx context.Context, id uint32) (*job.JobCategory, error)
	Create(ctx context.Context, entity *job.JobCategory) error
	Update(ctx context.Context, entity *job.JobCategory) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, uuid string) (*job.JobCategory, error)
}
