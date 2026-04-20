package user

import (
	"backend-app/internal/modules/master/model/user"
	"time"
)

type UserResponse struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FullName  string     `json:"full_name"`
	NIP       string     `json:"nip"`
	Role      string     `json:"role"`
	IsActive  bool       `json:"is_active"`
	LastLogin *time.Time `json:"last_login"`
	CreatedAt time.Time  `json:"created_at"`
}

func FromUser(u *user.User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FullName:  u.FullName,
		NIP:       u.NIP,
		Role:      u.Role,
		IsActive:  u.IsActive,
		LastLogin: u.LastLogin,
		CreatedAt: u.CreatedAt,
	}
}

func FromUsers(users []user.User) []UserResponse {
	var res []UserResponse
	for _, u := range users {
		res = append(res, *FromUser(&u))
	}
	return res
}
