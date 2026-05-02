package employment_status

type CreateEmploymentStatusRequest struct {
	Code           string `json:"code" binding:"required,max=20"`
	EmployeeStatus string `json:"employee_status" binding:"required,max=100"`
}

type UpdateEmploymentStatusRequest struct {
	Code           string `json:"code" binding:"required,max=20"`
	EmployeeStatus string `json:"employee_status" binding:"required,max=100"`
}
