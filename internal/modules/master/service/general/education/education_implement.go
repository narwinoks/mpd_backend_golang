package education

import (
	"backend-app/internal/core/exception"
	model "backend-app/internal/modules/master/model/general"
	repo "backend-app/internal/modules/master/repository/general/education"
	req "backend-app/internal/modules/master/request/general/education"
	res "backend-app/internal/modules/master/response/general/education"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type educationServiceImpl struct {
	repo repo.EducationRepository
}

func NewEducationService(repo repo.EducationRepository) EducationService {
	return &educationServiceImpl{
		repo: repo,
	}
}

func (s *educationServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.EducationResponse, *pagination.Meta, error) {
	educations, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch educations: %v", err)
		return nil, nil, err
	}

	var response []res.EducationResponse
	for _, e := range educations {
		response = append(response, res.EducationResponse{
			ID:            e.UUID,
			IsActive:      e.IsActive,
			EducationType: e.EducationType,
			Code:          e.Code,
			Name:          e.Name,
			SortOrder:     e.SortOrder,
			ExternalCode:  e.ExternalCode,
			CreatedAt:     e.CreatedAt,
			UpdatedAt:     e.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *educationServiceImpl) GetByID(ctx context.Context, id string) (*res.EducationResponse, error) {
	education, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Education not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.EducationResponse{
		ID:            education.UUID,
		EducationType: education.EducationType,
		Code:          education.Code,
		IsActive:      education.IsActive,
		Name:          education.Name,
		SortOrder:     education.SortOrder,
		ExternalCode:  education.ExternalCode,
		CreatedAt:     education.CreatedAt,
		UpdatedAt:     education.UpdatedAt,
	}, nil
}

func (s *educationServiceImpl) Create(ctx context.Context, request req.CreateEducationRequest) (*res.EducationResponse, error) {
	education := &model.Education{
		EducationType: request.EducationType,
		Code:          request.Code,
		Name:          request.Name,
		SortOrder:     request.SortOrder,
	}

	err := s.repo.Create(ctx, education)
	if err != nil {
		logrus.Errorf("Failed to create education: %v", err)
		return nil, err
	}

	return &res.EducationResponse{
		ID:            education.UUID,
		EducationType: education.EducationType,
		Code:          education.Code,
		Name:          education.Name,
		SortOrder:     education.SortOrder,
		ExternalCode:  education.ExternalCode,
		CreatedAt:     education.CreatedAt,
		UpdatedAt:     education.UpdatedAt,
	}, nil
}

func (s *educationServiceImpl) Update(ctx context.Context, id string, request req.UpdateEducationRequest) (*res.EducationResponse, error) {
	education, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Education not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	education.EducationType = request.EducationType
	education.Code = request.Code
	education.Name = request.Name
	education.SortOrder = request.SortOrder

	err = s.repo.Update(ctx, education)
	if err != nil {
		logrus.Errorf("Failed to update education: %v", err)
		return nil, err
	}

	return &res.EducationResponse{
		ID:            education.UUID,
		EducationType: education.EducationType,
		Code:          education.Code,
		Name:          education.Name,
		SortOrder:     education.SortOrder,
		ExternalCode:  education.ExternalCode,
		CreatedAt:     education.CreatedAt,
		UpdatedAt:     education.UpdatedAt,
	}, nil
}

func (s *educationServiceImpl) Delete(ctx context.Context, id string) error {
	education, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Education not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, education.ID)
	if err != nil {
		logrus.Errorf("Failed to delete education: %v", err)
		return err
	}

	return nil
}
