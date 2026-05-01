package job_category

import (
	"backend-app/internal/modules/master/model/job"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type jobCategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewJobCategoryRepository(db *gorm.DB) JobCategoryRepository {
	return &jobCategoryRepositoryImpl{db: db}
}

type JobCategoryWithCount struct {
	job.JobCategory
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *jobCategoryRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]job.JobCategory, int64, error) {
	var results []JobCategoryWithCount
	var jobCategories []job.JobCategory
	var total int64

	err := r.db.WithContext(ctx).Model(&job.JobCategory{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("job_category", "code")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			jobCategories = append(jobCategories, res.JobCategory)
		}
	}

	return jobCategories, total, nil
}

func (r *jobCategoryRepositoryImpl) FindByID(ctx context.Context, id uint32) (*job.JobCategory, error) {
	var entity job.JobCategory
	err := r.db.WithContext(ctx).Select("id", "uuid", "code", "job_category", "is_active", "created_at", "updated_at").First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *jobCategoryRepositoryImpl) Create(ctx context.Context, entity *job.JobCategory) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *jobCategoryRepositoryImpl) Update(ctx context.Context, entity *job.JobCategory) error {
	return r.db.WithContext(ctx).Updates(entity).Error
}

func (r *jobCategoryRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var entity job.JobCategory
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return err
	}
	return entity.SetNonActive(r.db.WithContext(ctx).Model(&entity))
}

func (r *jobCategoryRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*job.JobCategory, error) {
	var entity job.JobCategory
	err := r.db.WithContext(ctx).Select("id", "uuid", "code", "job_category", "is_active", "created_at", "updated_at").Where("uuid = ?", uuid).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
