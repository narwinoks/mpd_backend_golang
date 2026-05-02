package employment_status

import (
	"backend-app/internal/modules/master/model/job"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type employmentStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewEmploymentStatusRepository(db *gorm.DB) EmploymentStatusRepository {
	return &employmentStatusRepositoryImpl{db: db}
}

type EmploymentStatusWithCount struct {
	job.EmploymentStatus
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *employmentStatusRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]job.EmploymentStatus, int64, error) {
	var results []EmploymentStatusWithCount
	var entities []job.EmploymentStatus
	var total int64

	err := r.db.WithContext(ctx).Model(&job.EmploymentStatus{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("employee_status", "code")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			entities = append(entities, res.EmploymentStatus)
		}
	}

	return entities, total, nil
}

func (r *employmentStatusRepositoryImpl) FindByID(ctx context.Context, id uint32) (*job.EmploymentStatus, error) {
	var entity job.EmploymentStatus
	err := r.db.WithContext(ctx).Select("id", "uuid", "code", "employee_status", "is_active", "created_at", "updated_at").First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *employmentStatusRepositoryImpl) Create(ctx context.Context, entity *job.EmploymentStatus) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *employmentStatusRepositoryImpl) Update(ctx context.Context, entity *job.EmploymentStatus) error {
	return r.db.WithContext(ctx).Updates(entity).Error
}

func (r *employmentStatusRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var entity job.EmploymentStatus
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return err
	}
	return entity.SetNonActive(r.db.WithContext(ctx).Model(&entity))
}

func (r *employmentStatusRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*job.EmploymentStatus, error) {
	var entity job.EmploymentStatus
	err := r.db.WithContext(ctx).Select("id", "uuid", "code", "employee_status", "is_active", "created_at", "updated_at").Where("uuid = ?", uuid).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
