package role

type CreateRoleRequest struct {
	Role string `json:"role" binding:"required,max=100"`
}

type UpdateRoleRequest struct {
	Role string `json:"role" binding:"required,max=100"`
}
