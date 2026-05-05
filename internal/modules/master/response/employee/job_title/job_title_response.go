package job_title

import "time"

type JobTitleResponse struct {
	ID          string          `json:"id"`
	JobCategory JobCategoryInfo `json:"job_category"`
	Code        string          `json:"code"`
	JobTitle    string          `json:"job_title"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type JobCategoryInfo struct {
	ID          string `json:"id"`
	JobCategory string `json:"job_category"`
}
