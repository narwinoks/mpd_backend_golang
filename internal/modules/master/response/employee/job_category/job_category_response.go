package job_category

import "time"

type JobCategoryResponse struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	JobCategory string    `json:"job_category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
