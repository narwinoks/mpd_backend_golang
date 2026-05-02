package department

import (
	model "backend-app/internal/modules/master/model/department"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type departmentRepositoryImpl struct{ db *gorm.DB }

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepositoryImpl{db: db}
}

type departmentWithCount struct {
	model.Department
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *departmentRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]model.Department, int64, error) {
	var rows []departmentWithCount
	err := r.db.WithContext(ctx).Model(&model.Department{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("department_name", "external_code")).
		Find(&rows).Error
	if err != nil {
		return nil, 0, err
	}
	var results []model.Department
	var total int64
	for _, row := range rows {
		total = row.TotalCount
		results = append(results, row.Department)
	}
	return results, total, nil
}

func (r *departmentRepositoryImpl) FindByID(ctx context.Context, id uint32) (*model.Department, error) {
	var m model.Department
	err := r.db.WithContext(ctx).
		Select("id", "uuid", "department_name", "is_active", "external_code", "created_at", "updated_at").
		First(&m, id).Error
	return &m, err
}

func (r *departmentRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*model.Department, error) {
	var m model.Department
	err := r.db.WithContext(ctx).
		Select("id", "uuid", "department_name", "is_active", "external_code", "created_at", "updated_at").
		Where("uuid = ?", uuid).First(&m).Error
	return &m, err
}

func (r *departmentRepositoryImpl) Create(ctx context.Context, m *model.Department) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *departmentRepositoryImpl) Update(ctx context.Context, m *model.Department) error {
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *departmentRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var m model.Department
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return err
	}
	return m.SetNonActive(r.db.WithContext(ctx).Model(&m))
}
