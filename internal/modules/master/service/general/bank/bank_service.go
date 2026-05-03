package bank

import (
	req "backend-app/internal/modules/master/request/general/bank"
	res "backend-app/internal/modules/master/response/general/bank"
	"backend-app/pkg/pagination"
	"context"
)

type BankService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.BankResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.BankResponse, error)
	Create(ctx context.Context, request req.CreateBankRequest) (*res.BankResponse, error)
	Update(ctx context.Context, id string, request req.UpdateBankRequest) (*res.BankResponse, error)
	Delete(ctx context.Context, id string) error
}
