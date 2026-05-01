package religion

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/general"
	repo "backend-app/internal/modules/master/repository/general/religion"
	req "backend-app/internal/modules/master/request/general/religion"
	res "backend-app/internal/modules/master/response/general/religion"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type religionServiceImpl struct {
	repo repo.ReligionRepository
}

func NewReligionService(repo repo.ReligionRepository) ReligionService {
	return &religionServiceImpl{
		repo: repo,
	}
}

func (s *religionServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.ReligionResponse, *pagination.Meta, error) {
	religions, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch religions: %v", err)
		return nil, nil, err
	}

	var response []res.ReligionResponse
	for _, r := range religions {
		response = append(response, res.ReligionResponse{
			ID:        r.UUID,
			Religion:  r.Religion,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *religionServiceImpl) GetByID(ctx context.Context, id string) (*res.ReligionResponse, error) {
	rel, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Religion not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.ReligionResponse{
		ID:        rel.UUID,
		Religion:  rel.Religion,
		CreatedAt: rel.CreatedAt,
		UpdatedAt: rel.UpdatedAt,
	}, nil
}

func (s *religionServiceImpl) Create(ctx context.Context, request req.CreateReligionRequest) (*res.ReligionResponse, error) {
	rel := &model.Religion{
		Religion: request.Religion,
	}

	err := s.repo.Create(ctx, rel)
	if err != nil {
		logrus.Errorf("Failed to create religion: %v", err)
		return nil, err
	}

	return &res.ReligionResponse{
		ID:        rel.UUID,
		Religion:  rel.Religion,
		CreatedAt: rel.CreatedAt,
		UpdatedAt: rel.UpdatedAt,
	}, nil
}

func (s *religionServiceImpl) Update(ctx context.Context, id string, request req.UpdateReligionRequest) (*res.ReligionResponse, error) {
	rel, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Religion not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	rel.Religion = request.Religion

	err = s.repo.Update(ctx, rel)
	if err != nil {
		logrus.Errorf("Failed to update religion: %v", err)
		return nil, err
	}

	return &res.ReligionResponse{
		ID:        rel.UUID,
		Religion:  rel.Religion,
		CreatedAt: rel.CreatedAt,
		UpdatedAt: rel.UpdatedAt,
	}, nil
}

func (s *religionServiceImpl) Delete(ctx context.Context, id string) error {
	rel, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Religion not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, rel.ID)
	if err != nil {
		logrus.Errorf("Failed to delete religion: %v", err)
		return err
	}

	return nil
}
