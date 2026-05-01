package job_title

import (
	"backend-app/internal/modules/master/model/job"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type jobTitleRepositoryImpl struct {
	db *gorm.DB
}

func NewJobTitleRepository(db *gorm.DB) JobTitleRepository {
	return &jobTitleRepositoryImpl{db: db}
}

type JobTitleWithCount struct {
	job.JobTitle
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *jobTitleRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]job.JobTitle, int64, error) {
	var results []JobTitleWithCount
	var jobTitles []job.JobTitle
	var total int64

	err := r.db.WithContext(ctx).Model(&job.JobTitle{}).
		Preload("JobCategory").
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("job_title", "code")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			jobTitles = append(jobTitles, res.JobTitle)
		}
	}

	return jobTitles, total, nil
}

func (r *jobTitleRepositoryImpl) FindByID(ctx context.Context, id uint32) (*job.JobTitle, error) {
	var entity job.JobTitle
	err := r.db.WithContext(ctx).Preload("JobCategory").First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *jobTitleRepositoryImpl) Create(ctx context.Context, entity *job.JobTitle) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *jobTitleRepositoryImpl) Update(ctx context.Context, entity *job.JobTitle) error {
	return r.db.WithContext(ctx).Updates(entity).Error
}

func (r *jobTitleRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var entity job.JobTitle
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return err
	}
	return entity.SetNonActive(r.db.WithContext(ctx).Model(&entity))
}

func (r *jobTitleRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*job.JobTitle, error) {
	var entity job.JobTitle
	err := r.db.WithContext(ctx).Preload("JobCategory").Where("uuid = ?", uuid).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
