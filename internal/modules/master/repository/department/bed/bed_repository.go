package bed

import (
	model "backend-app/internal/modules/master/model/department"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type bedRepositoryImpl struct{ db *gorm.DB }

func NewBedRepository(db *gorm.DB) BedRepository {
	return &bedRepositoryImpl{db: db}
}

type bedWithCount struct {
	model.Bed
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *bedRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]model.Bed, int64, error) {
	var rows []bedWithCount
	err := r.db.WithContext(ctx).Model(&model.Bed{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("bed_number", "description", "external_code")).
		Find(&rows).Error
	if err != nil {
		return nil, 0, err
	}
	var results []model.Bed
	var total int64
	for _, row := range rows {
		total = row.TotalCount
		results = append(results, row.Bed)
	}
	return results, total, nil
}

func (r *bedRepositoryImpl) FindByID(ctx context.Context, id uint32) (*model.Bed, error) {
	var m model.Bed
	err := r.db.WithContext(ctx).
		Select("id", "uuid", "bed_number", "description", "room_id", "bed_status_id", "merged_bed_id", "is_active", "external_code", "created_at", "updated_at").
		First(&m, id).Error
	return &m, err
}

func (r *bedRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*model.Bed, error) {
	var m model.Bed
	err := r.db.WithContext(ctx).
		Select("id", "uuid", "bed_number", "description", "room_id", "bed_status_id", "merged_bed_id", "is_active", "external_code", "created_at", "updated_at").
		Where("uuid = ?", uuid).First(&m).Error
	return &m, err
}

func (r *bedRepositoryImpl) Create(ctx context.Context, m *model.Bed) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *bedRepositoryImpl) Update(ctx context.Context, m *model.Bed) error {
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *bedRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var m model.Bed
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return err
	}
	return m.SetNonActive(r.db.WithContext(ctx).Model(&m))
}
