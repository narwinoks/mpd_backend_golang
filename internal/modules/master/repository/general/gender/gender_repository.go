package gender

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type genderRepositoryImpl struct {
	db *gorm.DB
}

func NewGenderRepository(db *gorm.DB) GenderRepository {
	return &genderRepositoryImpl{db: db}
}

type GenderWithCount struct {
	general.Gender
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *genderRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]general.Gender, int64, error) {
	var results []GenderWithCount
	var genders []general.Gender
	var total int64

	err := r.db.WithContext(ctx).Model(&general.Gender{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("gender", "code")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			genders = append(genders, res.Gender)
		}
	}

	return genders, total, nil
}

func (r *genderRepositoryImpl) FindByID(ctx context.Context, id uint32) (*general.Gender, error) {
	var g general.Gender
	err := r.db.WithContext(ctx).Select("id", "uuid", "code", "gender", "is_active", "created_at", "updated_at").First(&g, id).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *genderRepositoryImpl) Create(ctx context.Context, g *general.Gender) error {
	return r.db.WithContext(ctx).Create(g).Error
}

func (r *genderRepositoryImpl) Update(ctx context.Context, g *general.Gender) error {
	return r.db.WithContext(ctx).Updates(g).Error
}

func (r *genderRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var g general.Gender
	if err := r.db.WithContext(ctx).First(&g, id).Error; err != nil {
		return err
	}
	return g.SetNonActive(r.db.WithContext(ctx).Model(&g))
}

func (r *genderRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*general.Gender, error) {
	var g general.Gender
	err := r.db.WithContext(ctx).Select("id", "uuid", "code", "gender", "is_active", "created_at", "updated_at").Where("uuid = ?", uuid).First(&g).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}
