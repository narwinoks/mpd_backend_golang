package position

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/job"
	repo "backend-app/internal/modules/master/repository/employee/position"
	req "backend-app/internal/modules/master/request/employee/position"
	res "backend-app/internal/modules/master/response/employee/position"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type positionServiceImpl struct {
	repo repo.PositionRepository
}

func NewPositionService(repo repo.PositionRepository) PositionService {
	return &positionServiceImpl{
		repo: repo,
	}
}

func (s *positionServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.PositionResponse, *pagination.Meta, error) {
	positions, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch positions: %v", err)
		return nil, nil, err
	}

	var response []res.PositionResponse
	for _, p := range positions {
		response = append(response, res.PositionResponse{
			ID:        p.UUID,
			Position:  p.Position,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *positionServiceImpl) GetByID(ctx context.Context, id string) (*res.PositionResponse, error) {
	position, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Position not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.PositionResponse{
		ID:        position.UUID,
		Position:  position.Position,
		CreatedAt: position.CreatedAt,
		UpdatedAt: position.UpdatedAt,
	}, nil
}

func (s *positionServiceImpl) Create(ctx context.Context, request req.CreatePositionRequest) (*res.PositionResponse, error) {
	position := &model.Position{
		Position: request.Position,
	}

	err := s.repo.Create(ctx, position)
	if err != nil {
		logrus.Errorf("Failed to create position: %v", err)
		return nil, err
	}

	return &res.PositionResponse{
		ID:        position.UUID,
		Position:  position.Position,
		CreatedAt: position.CreatedAt,
		UpdatedAt: position.UpdatedAt,
	}, nil
}

func (s *positionServiceImpl) Update(ctx context.Context, id string, request req.UpdatePositionRequest) (*res.PositionResponse, error) {
	position, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Position not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	position.Position = request.Position

	err = s.repo.Update(ctx, position)
	if err != nil {
		logrus.Errorf("Failed to update position: %v", err)
		return nil, err
	}

	return &res.PositionResponse{
		ID:        position.UUID,
		Position:  position.Position,
		CreatedAt: position.CreatedAt,
		UpdatedAt: position.UpdatedAt,
	}, nil
}

func (s *positionServiceImpl) Delete(ctx context.Context, id string) error {
	position, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Position not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, position.ID)
	if err != nil {
		logrus.Errorf("Failed to delete position: %v", err)
		return err
	}

	return nil
}
