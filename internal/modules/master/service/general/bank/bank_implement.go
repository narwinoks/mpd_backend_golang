package bank

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/general"
	repo "backend-app/internal/modules/master/repository/general/bank"
	req "backend-app/internal/modules/master/request/general/bank"
	res "backend-app/internal/modules/master/response/general/bank"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type bankServiceImpl struct {
	repo repo.BankRepository
}

func NewBankService(repo repo.BankRepository) BankService {
	return &bankServiceImpl{
		repo: repo,
	}
}

func (s *bankServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.BankResponse, *pagination.Meta, error) {
	banks, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch banks: %v", err)
		return nil, nil, err
	}

	var response []res.BankResponse
	for _, b := range banks {
		response = append(response, res.BankResponse{
			ID:        b.UUID,
			Bank:      b.Bank,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *bankServiceImpl) GetByID(ctx context.Context, id string) (*res.BankResponse, error) {
	bank, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Bank not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.BankResponse{
		ID:        bank.UUID,
		Bank:      bank.Bank,
		CreatedAt: bank.CreatedAt,
		UpdatedAt: bank.UpdatedAt,
	}, nil
}

func (s *bankServiceImpl) Create(ctx context.Context, request req.CreateBankRequest) (*res.BankResponse, error) {
	bank := &model.Bank{
		Bank: request.Bank,
	}

	err := s.repo.Create(ctx, bank)
	if err != nil {
		logrus.Errorf("Failed to create bank: %v", err)
		return nil, err
	}

	return &res.BankResponse{
		ID:        bank.UUID,
		Bank:      bank.Bank,
		CreatedAt: bank.CreatedAt,
		UpdatedAt: bank.UpdatedAt,
	}, nil
}

func (s *bankServiceImpl) Update(ctx context.Context, id string, request req.UpdateBankRequest) (*res.BankResponse, error) {
	bank, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Bank not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	bank.Bank = request.Bank

	err = s.repo.Update(ctx, bank)
	if err != nil {
		logrus.Errorf("Failed to update bank: %v", err)
		return nil, err
	}

	return &res.BankResponse{
		ID:        bank.UUID,
		Bank:      bank.Bank,
		CreatedAt: bank.CreatedAt,
		UpdatedAt: bank.UpdatedAt,
	}, nil
}

func (s *bankServiceImpl) Delete(ctx context.Context, id string) error {
	bank, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Bank not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, bank.ID)
	if err != nil {
		logrus.Errorf("Failed to delete bank: %v", err)
		return err
	}

	return nil
}
