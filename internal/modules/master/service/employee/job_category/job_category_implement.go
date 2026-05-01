package job_category

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/master/model/job"
	repo "backend-app/internal/modules/master/repository/employee/job_category"
	req "backend-app/internal/modules/master/request/employee/job_category"
	res "backend-app/internal/modules/master/response/employee/job_category"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type jobCategoryServiceImpl struct {
	repo repo.JobCategoryRepository
}

func NewJobCategoryService(repo repo.JobCategoryRepository) JobCategoryService {
	return &jobCategoryServiceImpl{
		repo: repo,
	}
}

func (s *jobCategoryServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.JobCategoryResponse, *pagination.Meta, error) {
	jobCategories, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch job categories: %v", err)
		return nil, nil, err
	}

	var response []res.JobCategoryResponse
	for _, j := range jobCategories {
		response = append(response, res.JobCategoryResponse{
			ID:          j.UUID,
			Code:        j.Code,
			JobCategory: j.JobCategory,
			CreatedAt:   j.CreatedAt,
			UpdatedAt:   j.UpdatedAt,
		})
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *jobCategoryServiceImpl) GetByID(ctx context.Context, id string) (*res.JobCategoryResponse, error) {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Job category not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return &res.JobCategoryResponse{
		ID:          entity.UUID,
		Code:        entity.Code,
		JobCategory: entity.JobCategory,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

func (s *jobCategoryServiceImpl) Create(ctx context.Context, request req.CreateJobCategoryRequest) (*res.JobCategoryResponse, error) {
	entity := &job.JobCategory{
		Code:        request.Code,
		JobCategory: request.JobCategory,
	}

	err := s.repo.Create(ctx, entity)
	if err != nil {
		logrus.Errorf("Failed to create job category: %v", err)
		return nil, err
	}

	return &res.JobCategoryResponse{
		ID:          entity.UUID,
		Code:        entity.Code,
		JobCategory: entity.JobCategory,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

func (s *jobCategoryServiceImpl) Update(ctx context.Context, id string, request req.UpdateJobCategoryRequest) (*res.JobCategoryResponse, error) {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Job category not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	entity.Code = request.Code
	entity.JobCategory = request.JobCategory

	err = s.repo.Update(ctx, entity)
	if err != nil {
		logrus.Errorf("Failed to update job category: %v", err)
		return nil, err
	}

	return &res.JobCategoryResponse{
		ID:          entity.UUID,
		Code:        entity.Code,
		JobCategory: entity.JobCategory,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}, nil
}

func (s *jobCategoryServiceImpl) Delete(ctx context.Context, id string) error {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Job category not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, entity.ID)
	if err != nil {
		logrus.Errorf("Failed to delete job category: %v", err)
		return err
	}

	return nil
}
