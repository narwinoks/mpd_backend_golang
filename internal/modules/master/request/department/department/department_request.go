package department

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name" binding:"required,max=255"`
	ExternalCode   string `json:"external_code"   binding:"omitempty,max=15"`
}

type UpdateDepartmentRequest struct {
	DepartmentName string `json:"department_name" binding:"required,max=255"`
	ExternalCode   string `json:"external_code"   binding:"omitempty,max=15"`
	IsActive       *bool  `json:"is_active"`
}
