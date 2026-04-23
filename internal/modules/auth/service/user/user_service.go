package user

import (
	req "backend-app/internal/modules/auth/request/user"
	res "backend-app/internal/modules/auth/response/user"
)

type UserService interface {
	Login(request *req.LoginRequest) (*res.LoginResponse, error)
}
