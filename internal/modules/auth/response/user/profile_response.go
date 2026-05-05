package user

type ProfileResponse struct {
	ID       uint32        `json:"id"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Role     RoleInfo      `json:"role"`
	Employee *EmployeeInfo `json:"employee"`
}

type RoleInfo struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

type EmployeeInfo struct {
	ID             uint32 `json:"id"`
	FullName       string `json:"full_name"`
	NIP            string `json:"nip"`
	IdentityNumber string `json:"identity_number"`
}
