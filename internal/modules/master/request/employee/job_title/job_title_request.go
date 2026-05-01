package job_title

type CreateJobTitleRequest struct {
	JobCategoryID string `json:"job_category_id" binding:"required"`
	Code          string `json:"code" binding:"required,max=20"`
	JobTitle      string `json:"job_title" binding:"required,max=100"`
}

type UpdateJobTitleRequest struct {
	JobCategoryID string `json:"job_category_id" binding:"required"`
	Code          string `json:"code" binding:"required,max=20"`
	JobTitle      string `json:"job_title" binding:"required,max=100"`
}
