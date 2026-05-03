package education

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type educationRepositoryImpl struct {
	db *gorm.DB
}

func NewEducationRepository(db *gorm.DB) EducationRepository {
	return &educationRepositoryImpl{db: db}
}

type EducationWithCount struct {
	general.Education
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *educationRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]general.Education, int64, error) {
	var results []EducationWithCount
	var educations []general.Education
	var total int64

	err := r.db.WithContext(ctx).Model(&general.Education{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("education")).
		Order("sort_order ASC").
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			educations = append(educations, res.Education)
		}
	}

	return educations, total, nil
}

func (r *educationRepositoryImpl) FindByID(ctx context.Context, id uint32) (*general.Education, error) {
	var education general.Education
	err := r.db.WithContext(ctx).Select("id", "uuid", "education_type", "code", "name", "sort_order", "is_active", "created_at", "updated_at", "external_code").First(&education, id).Error
	if err != nil {
		return nil, err
	}
	return &education, nil
}

func (r *educationRepositoryImpl) Create(ctx context.Context, education *general.Education) error {
	return r.db.WithContext(ctx).Create(education).Error
}

func (r *educationRepositoryImpl) Update(ctx context.Context, education *general.Education) error {
	return r.db.WithContext(ctx).Updates(education).Error
}

func (r *educationRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var education general.Education
	if err := r.db.WithContext(ctx).First(&education, id).Error; err != nil {
		return err
	}
	return education.SetNonActive(r.db.WithContext(ctx).Model(&education))
}

func (r *educationRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*general.Education, error) {
	var education general.Education
	err := r.db.WithContext(ctx).Select("id", "uuid", "education_type", "code", "name", "sort_order", "is_active", "created_at", "updated_at", "external_code").Where("uuid = ?", Uuid).First(&education).Error
	if err != nil {
		return nil, err
	}
	return &education, nil
}
