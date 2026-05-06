package funding_source

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/funding_source"
	repo "backend-app/internal/modules/master/repository/funding_source"
	req "backend-app/internal/modules/master/request/funding_source"
	res "backend-app/internal/modules/master/response/funding_source"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type fundingSourceServiceImpl struct {
	repo repo.FundingSourceRepository
}

func NewFundingSourceService(repo repo.FundingSourceRepository) FundingSourceService {
	return &fundingSourceServiceImpl{
		repo: repo,
	}
}

func (s *fundingSourceServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.FundingSourceResponse, *pagination.Meta, error) {
	fundingSources, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch funding sources: %v", err)
		return nil, nil, err
	}

	var response []res.FundingSourceResponse
	for _, fs := range fundingSources {
		response = append(response, res.FundingSourceResponse{
			ID:            fs.UUID,
			FundingSource: fs.FundingSource,
			ExternalCode:  fs.ExternalCode,
			IsActive:      fs.IsActive,
			CreatedAt:     fs.CreatedAt,
			UpdatedAt:     fs.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *fundingSourceServiceImpl) GetByID(ctx context.Context, id string) (*res.FundingSourceResponse, error) {
	fs, err := s.repo.FindByUUID(ctx, id)
	if err != nil {
		logrus.Errorf("Funding source not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.FundingSourceResponse{
		ID:            fs.UUID,
		FundingSource: fs.FundingSource,
		ExternalCode:  fs.ExternalCode,
		IsActive:      fs.IsActive,
		CreatedAt:     fs.CreatedAt,
		UpdatedAt:     fs.UpdatedAt,
	}, nil
}

func (s *fundingSourceServiceImpl) Create(ctx context.Context, request req.CreateFundingSourceRequest) (*res.FundingSourceResponse, error) {
	fs := &model.FundingSource{
		FundingSource: request.FundingSource,
	}
	fs.ExternalCode = request.ExternalCode

	err := s.repo.Create(ctx, fs)
	if err != nil {
		logrus.Errorf("Failed to create funding source: %v", err)
		return nil, err
	}

	return &res.FundingSourceResponse{
		ID:            fs.UUID,
		FundingSource: fs.FundingSource,
		ExternalCode:  fs.ExternalCode,
		IsActive:      fs.IsActive,
		CreatedAt:     fs.CreatedAt,
		UpdatedAt:     fs.UpdatedAt,
	}, nil
}

func (s *fundingSourceServiceImpl) Update(ctx context.Context, id string, request req.UpdateFundingSourceRequest) (*res.FundingSourceResponse, error) {
	fs, err := s.repo.FindByUUID(ctx, id)
	if err != nil {
		logrus.Errorf("Funding source not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	fs.FundingSource = request.FundingSource
	fs.ExternalCode = request.ExternalCode
	if request.IsActive != nil {
		fs.IsActive = *request.IsActive
	}

	err = s.repo.Update(ctx, fs)
	if err != nil {
		logrus.Errorf("Failed to update funding source: %v", err)
		return nil, err
	}

	return &res.FundingSourceResponse{
		ID:            fs.UUID,
		FundingSource: fs.FundingSource,
		ExternalCode:  fs.ExternalCode,
		IsActive:      fs.IsActive,
		CreatedAt:     fs.CreatedAt,
		UpdatedAt:     fs.UpdatedAt,
	}, nil
}

func (s *fundingSourceServiceImpl) Delete(ctx context.Context, id string) error {
	fs, err := s.repo.FindByUUID(ctx, id)
	if err != nil {
		logrus.Errorf("Funding source not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}

	err = s.repo.Delete(ctx, fs.ID)
	if err != nil {
		logrus.Errorf("Failed to delete funding source: %v", err)
		return err
	}

	return nil
}
