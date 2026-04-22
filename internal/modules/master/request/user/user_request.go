package user

type UserCreateRequest struct {
	Username   string `json:"username" binding:"required,min=4,max=50"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	RoleID     uint32 `json:"role_id" binding:"required"`
	EmployeeID uint32 `json:"employee_id" binding:"required"`
}

type UserUpdateRequest struct {
	FullName   string `json:"full_name" binding:"max=100"`
	NIP        string `json:"nip" binding:"max=20"`
	Role       string `json:"role" binding:"omitempty,oneof=admin doctor nurse staff"`
	IsActive   *bool  `json:"is_active"`
	RoleID     uint32 `json:"role_id" binding:"required"`
	EmployeeID uint32 `json:"employee_id" binding:"required"`
}
