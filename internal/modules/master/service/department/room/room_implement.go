package room

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/department"
	repo "backend-app/internal/modules/master/repository/department/room"
	req "backend-app/internal/modules/master/request/department/room"
	res "backend-app/internal/modules/master/response/department/room"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type roomServiceImpl struct{ repo repo.RoomRepository }

func NewRoomService(repo repo.RoomRepository) RoomService {
	return &roomServiceImpl{repo: repo}
}

func toResponse(m *model.Room) *res.RoomResponse {
	return &res.RoomResponse{
		ID:                 m.UUID,
		RoomName:           m.RoomName,
		WardID:             m.WardID,
		ClassID:            m.ClassID,
		RsOnlineCode:       m.RsOnlineCode,
		BedCount:           m.BedCount,
		OccupiedRoomCount:  m.OccupiedRoomCount,
		AvailableRoomCount: m.AvailableRoomCount,
		ExternalCode:       m.ExternalCode,
		IsActive:           m.IsActive,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
	}
}

func (s *roomServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.RoomResponse, *pagination.Meta, error) {
	items, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch rooms: %v", err)
		return nil, nil, err
	}
	var responses []res.RoomResponse
	for _, item := range items {
		responses = append(responses, *toResponse(&item))
	}
	return responses, pagination.BuildMeta(total, request.Page, request.Paginate, len(responses)), nil
}

func (s *roomServiceImpl) GetByID(ctx context.Context, id string) (*res.RoomResponse, error) {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return nil, exception.NewNotFoundError("Data not found")
	}
	return toResponse(m), nil
}

func (s *roomServiceImpl) Create(ctx context.Context, request req.CreateRoomRequest) (*res.RoomResponse, error) {
	m := &model.Room{
		RoomName:           request.RoomName,
		WardID:             request.WardID,
		ClassID:            request.ClassID,
		RsOnlineCode:       request.RsOnlineCode,
		BedCount:           request.BedCount,
		OccupiedRoomCount:  request.OccupiedRoomCount,
		AvailableRoomCount: request.AvailableRoomCount,
	}
	m.ExternalCode = request.ExternalCode
	if err := s.repo.Create(ctx, m); err != nil {
		logrus.Errorf("Failed to create room: %v", err)
		return nil, err
	}
	return toResponse(m), nil
}

func (s *roomServiceImpl) Update(ctx context.Context, id string, request req.UpdateRoomRequest) (*res.RoomResponse, error) {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return nil, exception.NewNotFoundError("Data not found")
	}
	m.RoomName = request.RoomName
	m.WardID = request.WardID
	m.ClassID = request.ClassID
	m.RsOnlineCode = request.RsOnlineCode
	m.BedCount = request.BedCount
	m.OccupiedRoomCount = request.OccupiedRoomCount
	m.AvailableRoomCount = request.AvailableRoomCount
	m.ExternalCode = request.ExternalCode
	if request.IsActive != nil {
		m.IsActive = *request.IsActive
	}
	if err := s.repo.Update(ctx, m); err != nil {
		logrus.Errorf("Failed to update room: %v", err)
		return nil, err
	}
	return toResponse(m), nil
}

func (s *roomServiceImpl) Delete(ctx context.Context, id string) error {
	m, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		return exception.NewNotFoundError("Data not found")
	}
	if err := s.repo.Delete(ctx, m.ID); err != nil {
		logrus.Errorf("Failed to delete room: %v", err)
		return err
	}
	return nil
}
