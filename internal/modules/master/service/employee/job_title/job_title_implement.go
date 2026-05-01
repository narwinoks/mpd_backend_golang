package job_title

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/master/model/job"
	repoJobCat "backend-app/internal/modules/master/repository/employee/job_category"
	repo "backend-app/internal/modules/master/repository/employee/job_title"
	req "backend-app/internal/modules/master/request/employee/job_title"
	res "backend-app/internal/modules/master/response/employee/job_title"
	"backend-app/pkg/pagination"
	"context"

	"github.com/sirupsen/logrus"
)

type jobTitleServiceImpl struct {
	repo       repo.JobTitleRepository
	repoJobCat repoJobCat.JobCategoryRepository
}

func NewJobTitleService(repo repo.JobTitleRepository, repoJobCat repoJobCat.JobCategoryRepository) JobTitleService {
	return &jobTitleServiceImpl{
		repo:       repo,
		repoJobCat: repoJobCat,
	}
}

func (s *jobTitleServiceImpl) GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.JobTitleResponse, *pagination.Meta, error) {
	entities, total, err := s.repo.FindAll(ctx, request)
	if err != nil {
		logrus.Errorf("Failed to fetch job titles: %v", err)
		return nil, nil, err
	}

	var response []res.JobTitleResponse
	for _, e := range entities {
		response = append(response, *s.mapToResponse(&e))
	}

	meta := pagination.BuildMeta(total, request.Page, request.Paginate, len(response))

	return response, meta, nil
}

func (s *jobTitleServiceImpl) GetByID(ctx context.Context, id string) (*res.JobTitleResponse, error) {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Job title not found with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	return s.mapToResponse(entity), nil
}

func (s *jobTitleServiceImpl) Create(ctx context.Context, request req.CreateJobTitleRequest) (*res.JobTitleResponse, error) {
	jobCat, err := s.repoJobCat.FindByUuid(ctx, request.JobCategoryID)
	if err != nil {
		return nil, exception.NewBadRequestError("Invalid Job Category ID")
	}

	entity := &job.JobTitle{
		JobCategoryID: jobCat.ID,
		Code:          request.Code,
		JobTitle:      request.JobTitle,
	}

	err = s.repo.Create(ctx, entity)
	if err != nil {
		logrus.Errorf("Failed to create job title: %v", err)
		return nil, err
	}

	// Reload to get preloaded category
	entity, _ = s.repo.FindByID(ctx, entity.ID)

	return s.mapToResponse(entity), nil
}

func (s *jobTitleServiceImpl) Update(ctx context.Context, id string, request req.UpdateJobTitleRequest) (*res.JobTitleResponse, error) {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Job title not found for update with id %s: %v", id, err)
		return nil, exception.NewNotFoundError("Data not found")
	}

	jobCat, err := s.repoJobCat.FindByUuid(ctx, request.JobCategoryID)
	if err != nil {
		return nil, exception.NewBadRequestError("Invalid Job Category ID")
	}

	entity.JobCategoryID = jobCat.ID
	entity.Code = request.Code
	entity.JobTitle = request.JobTitle

	err = s.repo.Update(ctx, entity)
	if err != nil {
		logrus.Errorf("Failed to update job title: %v", err)
		return nil, err
	}

	// Reload to get preloaded category
	entity, _ = s.repo.FindByID(ctx, entity.ID)

	return s.mapToResponse(entity), nil
}

func (s *jobTitleServiceImpl) Delete(ctx context.Context, id string) error {
	entity, err := s.repo.FindByUuid(ctx, id)
	if err != nil {
		logrus.Errorf("Job title not found for deletion with id %s: %v", id, err)
		return exception.NewNotFoundError("Data not found")
	}
	err = s.repo.Delete(ctx, entity.ID)
	if err != nil {
		logrus.Errorf("Failed to delete job title: %v", err)
		return err
	}

	return nil
}

func (s *jobTitleServiceImpl) mapToResponse(e *job.JobTitle) *res.JobTitleResponse {
	response := &res.JobTitleResponse{
		ID:        e.UUID,
		Code:      e.Code,
		JobTitle:  e.JobTitle,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}

	if e.JobCategory.ID != 0 {
		response.JobCategory = res.JobCategoryInfo{
			ID:          e.JobCategory.UUID,
			JobCategory: e.JobCategory.JobCategory,
		}
	}

	return response
}
