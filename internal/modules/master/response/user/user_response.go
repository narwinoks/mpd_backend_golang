package user

import (
	"backend-app/internal/modules/auth/models"
	"time"
)

type UserResponse struct {
	ID        uint32                  `json:"id"`
	Username  string                  `json:"username"`
	Email     string                  `json:"email"`
	RoleID    uint32                  `json:"role_id"`
	IsActive  bool                    `json:"is_active"`
	CreatedAt time.Time               `json:"created_at"`
	Role      *RoleDetailResponse     `json:"role,omitempty"`
	Employee  *EmployeeDetailResponse `json:"employee,omitempty"`
}
type RoleDetailResponse struct {
	ID   uint32 `json:"id"`
	UUID string `json:"uuid"`
	Role string `json:"role"`
}
type EmployeeDetailResponse struct {
	ID       uint32 `json:"id"`
	UUID     string `json:"uuid"`
	FullName string `json:"full_name"`
}

func FromUser(u *models.User) *UserResponse {
	if u == nil {
		return nil
	}
	response := &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		RoleID:    u.RoleID,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
	}
	if u.Role.ID != 0 {
		response.Role = &RoleDetailResponse{
			ID:   u.Role.ID,
			UUID: u.Role.UUID,
			Role: u.Role.Role,
		}
	}
	if u.Employee != nil {
		response.Employee = &EmployeeDetailResponse{
			ID:       u.Employee.ID,
			UUID:     u.Employee.UUID,
			FullName: u.Employee.FullName,
		}
	}

	return response
}

func FromUsers(users []models.User) []UserResponse {
	var res []UserResponse
	for _, u := range users {
		res = append(res, *FromUser(&u))
	}
	return res
}
