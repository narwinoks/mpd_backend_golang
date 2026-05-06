package subdistrict

import (
	"backend-app/internal/core/database"
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/location"
	repo "backend-app/internal/modules/master/repository/location/subdistrict"
	req "backend-app/internal/modules/master/request/location/subdistrict"
	cityRes "backend-app/internal/modules/master/response/location/city"
	provinceRes "backend-app/internal/modules/master/response/location/province"
	res "backend-app/internal/modules/master/response/location/subdistrict"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type subdistrictServiceImpl struct {
	repo repo.SubdistrictRepository
	db   *gorm.DB
}

func NewSubdistrictService(repo repo.SubdistrictRepository, db *gorm.DB) SubdistrictService {
	return &subdistrictServiceImpl{
		repo: repo,
		db:   db,
	}
}

func (s *subdistrictServiceImpl) GetAll(ctx context.Context, request req.FindAllRequest) ([]res.SubdistrictResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch subdistricts: %v", err)
		return nil, nil, err
	}

	var response []res.SubdistrictResponse
	for _, item := range items {
		response = append(response, *s.mapToResponse(&item))
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *subdistrictServiceImpl) GetByID(ctx context.Context, id string) (*res.SubdistrictResponse, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Subdistrict not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return s.mapToResponse(item), nil
}

func (s *subdistrictServiceImpl) Create(ctx context.Context, request req.CreateSubdistrictRequest) (string, error) {
	provinceID, err := database.ResolveUUID(ctx, s.db, "provinces_m", request.ProvinceID)
	if err != nil {
		return "", exception.NewNotFoundError("Province not found")
	}

	cityID, err := database.ResolveUUID(ctx, s.db, "cities_m", request.CityID)
	if err != nil {
		return "", exception.NewNotFoundError("City not found")
	}

	item := &model.Subdistrict{
		ProvinceID:  provinceID,
		CityID:      cityID,
		Code:        request.Code,
		Subdistrict: request.Subdistrict,
	}

	err = s.repo.Create(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to create subdistrict: %v", err)
		return "", err
	}

	return item.UUID, nil
}

func (s *subdistrictServiceImpl) Update(ctx context.Context, id string, request req.UpdateSubdistrictRequest) (string, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Subdistrict not found for update with id %s: %v", id, err)
		return "", exception.NewNotFoundError("Data not found")
	}

	provinceID, err := database.ResolveUUID(ctx, s.db, "provinces_m", request.ProvinceID)
	if err != nil {
		return "", exception.NewNotFoundError("Province not found")
	}

	cityID, err := database.ResolveUUID(ctx, s.db, "cities_m", request.CityID)
	if err != nil {
		return "", exception.NewNotFoundError("City not found")
	}

	item.ProvinceID = provinceID
	item.CityID = cityID
	item.Code = request.Code
	item.Subdistrict = request.Subdistrict

	err = s.repo.Update(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to update subdistrict: %v", err)
		return "", err
	}

	return item.UUID, nil
}

func (s *subdistrictServiceImpl) Delete(ctx context.Context, id string) error {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Subdistrict not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, item.ID)
	if err != nil {
		logrus.Errorf("Failed to delete subdistrict: %v", err)
		return err
	}

	return nil
}

func (s *subdistrictServiceImpl) mapToResponse(item *model.Subdistrict) *res.SubdistrictResponse {
	response := &res.SubdistrictResponse{
		ID:          item.UUID,
		Code:        item.Code,
		Subdistrict: item.Subdistrict,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}

	if item.Province.ID != 0 {
		response.Province = &provinceRes.ProvinceResponse{
			ID:       item.Province.UUID,
			Code:     item.Province.Code,
			Province: item.Province.Province,
		}
	}

	if item.City.ID != 0 {
		response.City = &cityRes.CityResponse{
			ID:   item.City.UUID,
			Code: item.City.Code,
			City: item.City.City,
		}
	}

	return response
}
