package bed

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/department"
	repo "backend-app/internal/modules/master/repository/department/bed"
	req "backend-app/internal/modules/master/request/department/bed"
	res "backend-app/internal/modules/master/response/department/bed"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type bedServiceImpl struct{ repo repo.BedRepository }

func NewBedService(repo repo.BedRepository) BedService {
	return &bedServiceImpl{repo: repo}
}

func toResponse(m *model.Bed) *res.BedResponse {
	return &res.BedResponse{
		ID:           m.UUID,
		BedNumber:    m.BedNumber,
		Description:  m.Description,
		RoomID:       m.RoomID,
		BedStatusID:  m.BedStatusID,
		MergedBedID:  m.MergedBedID,
		ExternalCode: m.ExternalCode,
		IsActive:     m.IsActive,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func (s *bedServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.BedResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch beds: %v", err)
		return nil, nil, err
	}
	var responses []res.BedResponse
	for _, item := range items {
		responses = append(responses, *toResponse(&item))
	}
	return responses, pagination.BuildMeta(total, request.Page, request.Paginate, len(responses)), nil
}

func (s *bedServiceImpl) GetByID(ctx context.Context, id string) (*res.BedResponse, error) {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return nil, exception.NewNotFoundError("Data not found")
	}
	return toResponse(m), nil
}

func (s *bedServiceImpl) Create(ctx context.Context, request req.CreateBedRequest) (*res.BedResponse, error) {
	m := &model.Bed{
		BedNumber:   request.BedNumber,
		Description: request.Description,
		RoomID:      request.RoomID,
		BedStatusID: request.BedStatusID,
		MergedBedID: request.MergedBedID,
	}
	m.ExternalCode = request.ExternalCode
	if err := s.repo.Create(ctx, m); err != nil {
		logrus.Errorf("Failed to create bed: %v", err)
		return nil, err
	}
	return toResponse(m), nil
}

func (s *bedServiceImpl) Update(ctx context.Context, id string, request req.UpdateBedRequest) (*res.BedResponse, error) {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return nil, exception.NewNotFoundError("Data not found")
	}
	m.BedNumber = request.BedNumber
	m.Description = request.Description
	m.RoomID = request.RoomID
	m.BedStatusID = request.BedStatusID
	m.MergedBedID = request.MergedBedID
	m.ExternalCode = request.ExternalCode
	if request.IsActive != nil {
		m.IsActive = *request.IsActive
	}
	if err := s.repo.Update(ctx, m); err != nil {
		logrus.Errorf("Failed to update bed: %v", err)
		return nil, err
	}
	return toResponse(m), nil
}

func (s *bedServiceImpl) Delete(ctx context.Context, id string) error {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return exception.NewNotFoundError("Data not found")
	}
	if err := s.repo.Delete(ctx, m.ID); err != nil {
		logrus.Errorf("Failed to delete bed: %v", err)
		return err
	}
	return nil
}
