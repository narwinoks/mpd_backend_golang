package funding_source

import (
	model "backend-app/internal/modules/master/model/funding_source"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type fundingSourceRepositoryImpl struct {
	db *gorm.DB
}

func NewFundingSourceRepository(db *gorm.DB) FundingSourceRepository {
	return &fundingSourceRepositoryImpl{db: db}
}

type FundingSourceWithCount struct {
	model.FundingSource
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *fundingSourceRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]model.FundingSource, int64, error) {
	var results []FundingSourceWithCount
	var fundingSources []model.FundingSource
	var total int64

	err := r.db.WithContext(ctx).Model(&model.FundingSource{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("funding_source")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			fundingSources = append(fundingSources, res.FundingSource)
		}
	}

	return fundingSources, total, nil
}

func (r *fundingSourceRepositoryImpl) FindByID(ctx context.Context, id uint32) (*model.FundingSource, error) {
	var fs model.FundingSource
	err := r.db.WithContext(ctx).Select("id", "uuid", "funding_source", "external_code", "is_active", "created_at", "updated_at").First(&fs, id).Error
	if err != nil {
		return nil, err
	}
	return &fs, nil
}

func (r *fundingSourceRepositoryImpl) FindByUUID(ctx context.Context, uuid string) (*model.FundingSource, error) {
	var fs model.FundingSource
	err := r.db.WithContext(ctx).Select("id", "uuid", "funding_source", "external_code", "is_active", "created_at", "updated_at").Where("uuid = ?", uuid).First(&fs).Error
	if err != nil {
		return nil, err
	}
	return &fs, nil
}

func (r *fundingSourceRepositoryImpl) Create(ctx context.Context, fs *model.FundingSource) error {
	return r.db.WithContext(ctx).Create(fs).Error
}

func (r *fundingSourceRepositoryImpl) Update(ctx context.Context, fs *model.FundingSource) error {
	return r.db.WithContext(ctx).Updates(fs).Error
}

func (r *fundingSourceRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var fs model.FundingSource
	if err := r.db.WithContext(ctx).First(&fs, id).Error; err != nil {
		return err
	}
	return fs.SetNonActive(r.db.WithContext(ctx).Model(&fs))
}
