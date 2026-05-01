package job_title

import (
	"backend-app/internal/modules/master/model/job"
	"backend-app/pkg/pagination"
	"context"
)

type JobTitleRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]job.JobTitle, int64, error)
	FindByID(ctx context.Context, id uint32) (*job.JobTitle, error)
	Create(ctx context.Context, entity *job.JobTitle) error
	Update(ctx context.Context, entity *job.JobTitle) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, uuid string) (*job.JobTitle, error)
}
