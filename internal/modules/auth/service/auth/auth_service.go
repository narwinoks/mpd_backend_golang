package auth

import (
	req "backend-app/internal/modules/auth/request/user"
	res "backend-app/internal/modules/auth/response/user"
)

type AuthService interface {
	Login(request *req.LoginRequest) (*res.LoginResponse, error)
	RefreshToken(request *req.RefreshTokenRequest) (*res.LoginResponse, error)
	Logout(userID uint32) error
}
