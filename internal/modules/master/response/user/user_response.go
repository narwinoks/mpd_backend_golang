package user

import (
	"backend-app/internal/modules/auth/models"
	"time"
)

type UserResponse struct {
	ID        uint32    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	RoleID    uint32    `json:"role_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func FromUser(u *models.User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		RoleID:    u.RoleID,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
	}
}

func FromUsers(users []models.User) []UserResponse {
	var res []UserResponse
	for _, u := range users {
		res = append(res, *FromUser(&u))
	}
	return res
}
