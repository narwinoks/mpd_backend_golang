package job_category

type CreateJobCategoryRequest struct {
	Code        string `json:"code" binding:"required,max=20"`
	JobCategory string `json:"job_category" binding:"required,max=100"`
}

type UpdateJobCategoryRequest struct {
	Code        string `json:"code" binding:"required,max=20"`
	JobCategory string `json:"job_category" binding:"required,max=100"`
}
