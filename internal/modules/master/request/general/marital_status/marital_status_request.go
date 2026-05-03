package marital_status

type CreateMaritalStatusRequest struct {
	MaritalStatus string `json:"material_status" binding:"required,max=50"`
}

type UpdateMaritalStatusRequest struct {
	MaritalStatus string `json:"material_status" binding:"required,max=50"`
}
