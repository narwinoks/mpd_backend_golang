package bank

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type bankRepositoryImpl struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepositoryImpl{db: db}
}

type BankWithCount struct {
	general.Bank
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *bankRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]general.Bank, int64, error) {
	var results []BankWithCount
	var banks []general.Bank
	var total int64

	err := r.db.WithContext(ctx).Model(&general.Bank{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("bank")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			banks = append(banks, res.Bank)
		}
	}

	return banks, total, nil
}

func (r *bankRepositoryImpl) FindByID(ctx context.Context, id uint32) (*general.Bank, error) {
	var bank general.Bank
	err := r.db.WithContext(ctx).Select("id", "uuid", "bank", "is_active", "created_at", "updated_at").First(&bank, id).Error
	if err != nil {
		return nil, err
	}
	return &bank, nil
}

func (r *bankRepositoryImpl) Create(ctx context.Context, bank *general.Bank) error {
	return r.db.WithContext(ctx).Create(bank).Error
}

func (r *bankRepositoryImpl) Update(ctx context.Context, bank *general.Bank) error {
	return r.db.WithContext(ctx).Updates(bank).Error
}

func (r *bankRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var bank general.Bank
	if err := r.db.WithContext(ctx).First(&bank, id).Error; err != nil {
		return err
	}
	return bank.SetNonActive(r.db.WithContext(ctx).Model(&bank))
}

func (r *bankRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*general.Bank, error) {
	var bank general.Bank
	err := r.db.WithContext(ctx).Select("id", "uuid", "bank", "is_active", "created_at", "updated_at").Where("uuid = ?", Uuid).First(&bank).Error
	if err != nil {
		return nil, err
	}
	return &bank, nil
}
