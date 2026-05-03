package marital_status

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type maritalStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewMaritalStatusRepository(db *gorm.DB) MaritalStatusRepository {
	return &maritalStatusRepositoryImpl{db: db}
}

type MaritalStatusWithCount struct {
	general.MaritalStatus
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *maritalStatusRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]general.MaritalStatus, int64, error) {
	var results []MaritalStatusWithCount
	var items []general.MaritalStatus
	var total int64

	err := r.db.WithContext(ctx).Model(&general.MaritalStatus{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("marital_status")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			items = append(items, res.MaritalStatus)
		}
	}

	return items, total, nil
}

func (r *maritalStatusRepositoryImpl) FindByID(ctx context.Context, id uint32) (*general.MaritalStatus, error) {
	var maritalStatus general.MaritalStatus
	err := r.db.WithContext(ctx).Select("id", "uuid", "material_status", "is_active", "created_at", "updated_at").First(&maritalStatus, id).Error
	if err != nil {
		return nil, err
	}
	return &maritalStatus, nil
}

func (r *maritalStatusRepositoryImpl) Create(ctx context.Context, maritalStatus *general.MaritalStatus) error {
	return r.db.WithContext(ctx).Create(maritalStatus).Error
}

func (r *maritalStatusRepositoryImpl) Update(ctx context.Context, maritalStatus *general.MaritalStatus) error {
	return r.db.WithContext(ctx).Updates(maritalStatus).Error
}

func (r *maritalStatusRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var maritalStatus general.MaritalStatus
	if err := r.db.WithContext(ctx).First(&maritalStatus, id).Error; err != nil {
		return err
	}
	return maritalStatus.SetNonActive(r.db.WithContext(ctx).Model(&maritalStatus))
}

func (r *maritalStatusRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*general.MaritalStatus, error) {
	var maritalStatus general.MaritalStatus
	err := r.db.WithContext(ctx).Select("id", "uuid", "material_status", "is_active", "created_at", "updated_at").Where("uuid = ?", Uuid).First(&maritalStatus).Error
	if err != nil {
		return nil, err
	}
	return &maritalStatus, nil
}
