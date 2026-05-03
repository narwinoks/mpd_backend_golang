package education

import (
	"time"
)

type EducationResponse struct {
	ID            string    `json:"id"`
	EducationType string    `json:"education_type"`
	IsActive      bool      `json:"is_active"`
	Code          string    `json:"code"`
	Name          string    `json:"name"`
	SortOrder     int       `json:"sort_order"`
	ExternalCode  string    `json:"external_code"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
