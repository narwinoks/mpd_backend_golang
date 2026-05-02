package ward

type CreateWardRequest struct {
	WardName          string  `json:"ward_name"           binding:"required,max=255"`
	DepartmentID      *uint32 `json:"department_id"`
	IsExecutive       bool    `json:"is_executive"`
	Icon              *string `json:"icon"                binding:"omitempty,max=100"`
	QueueNumberPrefix *string `json:"queue_number_prefix" binding:"omitempty,max=200"`
	ExternalCode      string  `json:"external_code"       binding:"omitempty,max=20"`
}

type UpdateWardRequest struct {
	WardName          string  `json:"ward_name"           binding:"required,max=255"`
	DepartmentID      *uint32 `json:"department_id"`
	IsExecutive       bool    `json:"is_executive"`
	Icon              *string `json:"icon"                binding:"omitempty,max=100"`
	QueueNumberPrefix *string `json:"queue_number_prefix" binding:"omitempty,max=200"`
	ExternalCode      string  `json:"external_code"       binding:"omitempty,max=20"`
	IsActive          *bool   `json:"is_active"`
}
