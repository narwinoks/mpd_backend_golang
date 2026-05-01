package user

import (
	req "backend-app/internal/modules/master/request/user"
	res "backend-app/internal/modules/master/response/user"
	"backend-app/pkg/pagination"
	"context"
)

type UserService interface {
	GetAllUsers(ctx context.Context, req pagination.BaseRequest) ([]res.UserResponse, *pagination.Meta, error)
	GetUserByID(ctx context.Context, id uint) (*res.UserResponse, error)
	CreateUser(ctx context.Context, req *req.UserCreateRequest) (*res.UserResponse, error)
}
