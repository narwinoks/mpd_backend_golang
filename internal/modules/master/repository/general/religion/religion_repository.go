package religion

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type religionRepositoryImpl struct {
	db *gorm.DB
}

func NewReligionRepository(db *gorm.DB) ReligionRepository {
	return &religionRepositoryImpl{db: db}
}

type ReligionWithCount struct {
	general.Religion
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *religionRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]general.Religion, int64, error) {
	var results []ReligionWithCount
	var religions []general.Religion
	var total int64

	err := r.db.WithContext(ctx).Model(&general.Religion{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("religion")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			religions = append(religions, res.Religion)
		}
	}

	return religions, total, nil
}

func (r *religionRepositoryImpl) FindByID(ctx context.Context, id uint32) (*general.Religion, error) {
	var rel general.Religion
	err := r.db.WithContext(ctx).Select("id", "uuid", "religion", "is_active", "created_at", "updated_at").First(&rel, id).Error
	if err != nil {
		return nil, err
	}
	return &rel, nil
}

func (r *religionRepositoryImpl) Create(ctx context.Context, rel *general.Religion) error {
	return r.db.WithContext(ctx).Create(rel).Error
}

func (r *religionRepositoryImpl) Update(ctx context.Context, rel *general.Religion) error {
	return r.db.WithContext(ctx).Updates(rel).Error
}

func (r *religionRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var rel general.Religion
	if err := r.db.WithContext(ctx).First(&rel, id).Error; err != nil {
		return err
	}
	return rel.SetNonActive(r.db.WithContext(ctx).Model(&rel))
}

func (r *religionRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*general.Religion, error) {
	var rel general.Religion
	err := r.db.WithContext(ctx).Select("id", "uuid", "religion", "is_active", "created_at", "updated_at").Where("uuid = ?", uuid).First(&rel).Error
	if err != nil {
		return nil, err
	}
	return &rel, nil
}
