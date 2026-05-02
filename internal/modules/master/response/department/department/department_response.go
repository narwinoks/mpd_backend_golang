package department

import "time"

type DepartmentResponse struct {
	ID             string    `json:"id"`
	DepartmentName string    `json:"department_name"`
	ExternalCode   string    `json:"external_code"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
