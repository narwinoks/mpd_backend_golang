package user

import (
	req "backend-app/internal/modules/master/request/user"
	res "backend-app/internal/modules/master/response/user"
	"context"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]res.UserResponse, error)
	GetUserByID(ctx context.Context, id uint) (*res.UserResponse, error)
	CreateUser(ctx context.Context, req *req.UserCreateRequest) (*res.UserResponse, error)
}
