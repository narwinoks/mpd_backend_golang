package user

type UserCreateRequest struct {
	Username string `json:"username" binding:"required,min=4,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required,max=100"`
	NIP      string `json:"nip" binding:"max=20"`
	Role     string `json:"role" binding:"required,oneof=admin doctor nurse staff"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name" binding:"max=100"`
	NIP      string `json:"nip" binding:"max=20"`
	Role     string `json:"role" binding:"omitempty,oneof=admin doctor nurse staff"`
	IsActive *bool  `json:"is_active"`
}
