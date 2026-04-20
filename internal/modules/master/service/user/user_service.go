package user

import (
	req "backend-app/internal/modules/master/request/user"
	res "backend-app/internal/modules/master/response/user"
)

type UserService interface {
	GetAllUsers() ([]res.UserResponse, error)
	GetUserByID(id uint) (*res.UserResponse, error)
	CreateUser(req *req.UserCreateRequest) (*res.UserResponse, error)
}
