package city

import (
	"backend-app/internal/core/database"
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/location"
	repo "backend-app/internal/modules/master/repository/location/city"
	req "backend-app/internal/modules/master/request/location/city"
	res "backend-app/internal/modules/master/response/location/city"
	provinceRes "backend-app/internal/modules/master/response/location/province"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type cityServiceImpl struct {
	repo repo.CityRepository
	db   *gorm.DB
}

func NewCityService(repo repo.CityRepository, db *gorm.DB) CityService {
	return &cityServiceImpl{
		repo: repo,
		db:   db,
	}
}

func (s *cityServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.CityResponse, *pagination.Meta, error) {
	cities, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch cities: %v", err)
		return nil, nil, err
	}

	var response []res.CityResponse
	for _, c := range cities {
		response = append(response, *s.mapToResponse(&c))
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *cityServiceImpl) GetByID(ctx context.Context, id string) (*res.CityResponse, error) {
	city, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("City not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return s.mapToResponse(city), nil
}

func (s *cityServiceImpl) Create(ctx context.Context, request req.CreateCityRequest) (string, error) {
	provinceID, err := database.ResolveUUID(ctx, s.db, "provinces_m", request.ProvinceID)
	if err != nil {
		return "", exception.NewNotFoundError("Province not found")
	}

	city := &model.City{
		ProvinceID: provinceID,
		Code:       request.Code,
		City:       request.City,
	}

	err = s.repo.Create(ctx, city)
	if err != nil {
		logrus.Errorf("Failed to create city: %v", err)
		return "", err
	}

	return city.UUID, nil
}

func (s *cityServiceImpl) Update(ctx context.Context, id string, request req.UpdateCityRequest) (string, error) {
	city, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("City not found for update with id %s: %v", id, err)
		return "", exception.NewNotFoundError("Data not found")
	}

	provinceID, err := database.ResolveUUID(ctx, s.db, "provinces_m", request.ProvinceID)
	if err != nil {
		return "", exception.NewNotFoundError("Province not found")
	}

	city.ProvinceID = provinceID
	city.Code = request.Code
	city.City = request.City

	err = s.repo.Update(ctx, city)
	if err != nil {
		logrus.Errorf("Failed to update city: %v", err)
		return "", err
	}

	return city.UUID, nil
}

func (s *cityServiceImpl) Delete(ctx context.Context, id string) error {
	city, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("City not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, city.ID)
	if err != nil {
		logrus.Errorf("Failed to delete city: %v", err)
		return err
	}

	return nil
}

func (s *cityServiceImpl) mapToResponse(c *model.City) *res.CityResponse {
	response := &res.CityResponse{
		ID:        c.UUID,
		Code:      c.Code,
		City:      c.City,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}

	if c.Province.ID != 0 {
		response.Province = &provinceRes.ProvinceResponse{
			ID:       c.Province.UUID,
			Code:     c.Province.Code,
			Province: c.Province.Province,
		}
	}

	return response
}
