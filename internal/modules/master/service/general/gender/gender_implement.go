package gender

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/general"
	repo "backend-app/internal/modules/master/repository/general/gender"
	req "backend-app/internal/modules/master/request/general/gender"
	res "backend-app/internal/modules/master/response/general/gender"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type genderServiceImpl struct {
	repo repo.GenderRepository
}

func NewGenderService(repo repo.GenderRepository) GenderService {
	return &genderServiceImpl{
		repo: repo,
	}
}

func (s *genderServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.GenderResponse, *pagination.Meta, error) {
	genders, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch genders: %v", err)
		return nil, nil, err
	}

	var response []res.GenderResponse
	for _, g := range genders {
		response = append(response, res.GenderResponse{
			ID:        g.UUID,
			Code:      g.Code,
			Gender:    g.Gender,
			CreatedAt: g.CreatedAt,
			UpdatedAt: g.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *genderServiceImpl) GetByID(ctx context.Context, id string) (*res.GenderResponse, error) {
	g, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Gender not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.GenderResponse{
		ID:        g.UUID,
		Code:      g.Code,
		Gender:    g.Gender,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}

func (s *genderServiceImpl) Create(ctx context.Context, request req.CreateGenderRequest) (*res.GenderResponse, error) {
	g := &model.Gender{
		Code:   request.Code,
		Gender: request.Gender,
	}

	err := s.repo.Create(ctx, g)
	if err != nil {
		logrus.Errorf("Failed to create gender: %v", err)
		return nil, err
	}

	return &res.GenderResponse{
		ID:        g.UUID,
		Code:      g.Code,
		Gender:    g.Gender,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}

func (s *genderServiceImpl) Update(ctx context.Context, id string, request req.UpdateGenderRequest) (*res.GenderResponse, error) {
	g, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Gender not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	g.Code = request.Code
	g.Gender = request.Gender

	err = s.repo.Update(ctx, g)
	if err != nil {
		logrus.Errorf("Failed to update gender: %v", err)
		return nil, err
	}

	return &res.GenderResponse{
		ID:        g.UUID,
		Code:      g.Code,
		Gender:    g.Gender,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}, nil
}

func (s *genderServiceImpl) Delete(ctx context.Context, id string) error {
	g, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Gender not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, g.ID)
	if err != nil {
		logrus.Errorf("Failed to delete gender: %v", err)
		return err
	}

	return nil
}
