package employment_status

import "time"

type EmploymentStatusResponse struct {
	ID             string    `json:"id"`
	Code           string    `json:"code"`
	EmployeeStatus string    `json:"employee_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
