package village

import (
	"backend-app/internal/core/database"
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/location"
	repo "backend-app/internal/modules/master/repository/location/village"
	req "backend-app/internal/modules/master/request/location/village"
	res "backend-app/internal/modules/master/response/location/village"
	cityRes "backend-app/internal/modules/master/response/location/city"
	provinceRes "backend-app/internal/modules/master/response/location/province"
	subdistrictRes "backend-app/internal/modules/master/response/location/subdistrict"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type villageServiceImpl struct {
	repo repo.VillageRepository
	db   *gorm.DB
}

func NewVillageService(repo repo.VillageRepository, db *gorm.DB) VillageService {
	return &villageServiceImpl{
		repo: repo,
		db:   db,
	}
}

func (s *villageServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.VillageResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch villages: %v", err)
		return nil, nil, err
	}

	var response []res.VillageResponse
	for _, item := range items {
		response = append(response, *s.mapToResponse(&item))
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *villageServiceImpl) GetByID(ctx context.Context, id string) (*res.VillageResponse, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Village not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return s.mapToResponse(item), nil
}

func (s *villageServiceImpl) Create(ctx context.Context, request req.CreateVillageRequest) (string, error) {
	provinceID, err := database.ResolveUUID(ctx, s.db, "provinces_m", request.ProvinceID)
	if err != nil {
		return "", exception.NewNotFoundError("Province not found")
	}

	cityID, err := database.ResolveUUID(ctx, s.db, "cities_m", request.CityID)
	if err != nil {
		return "", exception.NewNotFoundError("City not found")
	}

	subdistrictID, err := database.ResolveUUID(ctx, s.db, "subdistrict_m", request.SubdistrictID)
	if err != nil {
		return "", exception.NewNotFoundError("Subdistrict not found")
	}

	item := &model.Village{
		ProvinceID:    provinceID,
		CityID:        cityID,
		SubdistrictID: subdistrictID,
		Village:       request.Village,
		PostalCode:    request.PostalCode,
		Longitude:     request.Longitude,
		Latitude:      request.Latitude,
	}

	err = s.repo.Create(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to create village: %v", err)
		return "", err
	}

	return item.UUID, nil
}

func (s *villageServiceImpl) Update(ctx context.Context, id string, request req.UpdateVillageRequest) (string, error) {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Village not found for update with id %s: %v", id, err)
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

	subdistrictID, err := database.ResolveUUID(ctx, s.db, "subdistrict_m", request.SubdistrictID)
	if err != nil {
		return "", exception.NewNotFoundError("Subdistrict not found")
	}

	item.ProvinceID = provinceID
	item.CityID = cityID
	item.SubdistrictID = subdistrictID
	item.Village = request.Village
	item.PostalCode = request.PostalCode
	item.Longitude = request.Longitude
	item.Latitude = request.Latitude

	err = s.repo.Update(ctx, item)
	if err != nil {
		logrus.Errorf("Failed to update village: %v", err)
		return "", err
	}

	return item.UUID, nil
}

func (s *villageServiceImpl) Delete(ctx context.Context, id string) error {
	item, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Village not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, item.ID)
	if err != nil {
		logrus.Errorf("Failed to delete village: %v", err)
		return err
	}

	return nil
}

func (s *villageServiceImpl) mapToResponse(item *model.Village) *res.VillageResponse {
	response := &res.VillageResponse{
		ID:         item.UUID,
		Village:    item.Village,
		PostalCode: item.PostalCode,
		Longitude:  item.Longitude,
		Latitude:   item.Latitude,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}

	if item.Province.ID != 0 {
		response.Province = &provinceRes.ProvinceResponse{
			ID:       item.Province.UUID,
			Province: item.Province.Province,
		}
	}

	if item.City.ID != 0 {
		response.City = &cityRes.CityResponse{
			ID:   item.City.UUID,
			City: item.City.City,
		}
	}

	if item.Subdistrict.ID != 0 {
		response.Subdistrict = &subdistrictRes.SubdistrictResponse{
			ID:          item.Subdistrict.UUID,
			Subdistrict: item.Subdistrict.Subdistrict,
		}
	}

	return response
}
