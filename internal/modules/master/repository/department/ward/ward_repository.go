package ward

import (
	model "backend-app/internal/modules/master/model/department"
	"backend-app/pkg/pagination"
	"context"
	"strings"

	"gorm.io/gorm"
)

type wardRepositoryImpl struct{ db *gorm.DB }

func NewWardRepository(db *gorm.DB) WardRepository {
	return &wardRepositoryImpl{db: db}
}

type wardRow struct {
	model.Ward
	DepartmentName string `gorm:"column:department_name"`
}

func (r *wardRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]model.Ward, int64, error) {
	var total int64
	var rows []wardRow

	base := r.db.WithContext(ctx).
		Table("wards_m w").
		Joins("LEFT JOIN departments_m d ON d.id = w.department_id AND d.deleted_at IS NULL").
		Where("w.deleted_at IS NULL")

	search := strings.TrimSpace(strings.Trim(req.Search, "'\""))
	if search != "" {
		base = base.Where("(w.ward_name ILIKE ? OR w.external_code ILIKE ?)", "%"+search+"%", "%"+search+"%")
	}

	base.Count(&total)

	offset := (req.Page - 1) * req.Paginate
	if offset < 0 {
		offset = 0
	}
	limit := req.Paginate
	if limit <= 0 {
		limit = 10
	}

	err := base.
		Select("w.*, d.department_name").
		Offset(offset).Limit(limit).
		Scan(&rows).Error
	if err != nil {
		return nil, 0, err
	}

	results := make([]model.Ward, 0, len(rows))
	for _, row := range rows {
		w := row.Ward
		if row.DepartmentName != "" {
			w.Department = model.Department{DepartmentName: row.DepartmentName}
			if w.DepartmentID != nil {
				w.Department.BaseModel.ID = *w.DepartmentID
			}
		}
		results = append(results, w)
	}

	return results, total, nil
}

func (r *wardRepositoryImpl) FindByID(ctx context.Context, id uint32) (*model.Ward, error) {
	var m model.Ward
	err := r.db.WithContext(ctx).
		Preload("Department", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "department_name")
		}).
		First(&m, id).Error
	return &m, err
}

func (r *wardRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*model.Ward, error) {
	var m model.Ward
	err := r.db.WithContext(ctx).
		Preload("Department", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "department_name")
		}).
		Where("uuid = ?", uuid).First(&m).Error
	return &m, err
}

func (r *wardRepositoryImpl) Create(ctx context.Context, m *model.Ward) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *wardRepositoryImpl) Update(ctx context.Context, m *model.Ward) error {
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *wardRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var m model.Ward
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return err
	}
	return m.SetNonActive(r.db.WithContext(ctx).Model(&m))
}
