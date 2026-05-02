package employment_status

import (
	"backend-app/internal/modules/master/model/job"
	"backend-app/pkg/pagination"
	"context"
)

type EmploymentStatusRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]job.EmploymentStatus, int64, error)
	FindByID(ctx context.Context, id uint32) (*job.EmploymentStatus, error)
	Create(ctx context.Context, entity *job.EmploymentStatus) error
	Update(ctx context.Context, entity *job.EmploymentStatus) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, uuid string) (*job.EmploymentStatus, error)
}
