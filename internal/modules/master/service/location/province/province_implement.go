package province

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/location"
	repo "backend-app/internal/modules/master/repository/location/province"
	req "backend-app/internal/modules/master/request/location/province"
	res "backend-app/internal/modules/master/response/location/province"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type provinceServiceImpl struct {
	repo repo.ProvinceRepository
}

func NewProvinceService(repo repo.ProvinceRepository) ProvinceService {
	return &provinceServiceImpl{
		repo: repo,
	}
}

func (s *provinceServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.ProvinceResponse, *pagination.Meta, error) {
	provinces, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch provinces: %v", err)
		return nil, nil, err
	}

	var response []res.ProvinceResponse
	for _, p := range provinces {
		response = append(response, res.ProvinceResponse{
			ID:        p.UUID,
			Province:  p.Province,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *provinceServiceImpl) GetByID(ctx context.Context, id string) (*res.ProvinceResponse, error) {
	province, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Province not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.ProvinceResponse{
		ID:        province.UUID,
		Province:  province.Province,
		CreatedAt: province.CreatedAt,
		UpdatedAt: province.UpdatedAt,
	}, nil
}

func (s *provinceServiceImpl) Create(ctx context.Context, request req.CreateProvinceRequest) (*res.ProvinceResponse, error) {
	province := &model.Province{
		Province: request.Province,
	}

	err := s.repo.Create(ctx, province)
	if err != nil {
		logrus.Errorf("Failed to create province: %v", err)
		return nil, err
	}

	return &res.ProvinceResponse{
		ID:        province.UUID,
		Province:  province.Province,
		CreatedAt: province.CreatedAt,
		UpdatedAt: province.UpdatedAt,
	}, nil
}

func (s *provinceServiceImpl) Update(ctx context.Context, id string, request req.UpdateProvinceRequest) (*res.ProvinceResponse, error) {
	province, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Province not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	province.Province = request.Province

	err = s.repo.Update(ctx, province)
	if err != nil {
		logrus.Errorf("Failed to update province: %v", err)
		return nil, err
	}

	return &res.ProvinceResponse{
		ID:        province.UUID,
		Province:  province.Province,
		CreatedAt: province.CreatedAt,
		UpdatedAt: province.UpdatedAt,
	}, nil
}

func (s *provinceServiceImpl) Delete(ctx context.Context, id string) error {
	province, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Province not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, province.ID)
	if err != nil {
		logrus.Errorf("Failed to delete province: %v", err)
		return err
	}

	return nil
}
