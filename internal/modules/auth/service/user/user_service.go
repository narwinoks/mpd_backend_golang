package user

import (
	res "backend-app/internal/modules/auth/response/user"
)

type UserService interface {
	GetProfile(userID uint32) (*res.ProfileResponse, error)
}
