package bank

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"
)

type BankRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]general.Bank, int64, error)
	FindByID(ctx context.Context, id uint32) (*general.Bank, error)
	Create(ctx context.Context, bank *general.Bank) error
	Update(ctx context.Context, bank *general.Bank) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*general.Bank, error)
}
