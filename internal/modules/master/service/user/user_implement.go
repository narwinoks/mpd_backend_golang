package user

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/master/model"
	repo "backend-app/internal/modules/master/repository/user"
	req "backend-app/internal/modules/master/request/user"
	res "backend-app/internal/modules/master/response/user"
)

type userServiceImpl struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) GetAllUsers() ([]res.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// Requirement: Treat empty array as 404
	if len(users) == 0 {
		return nil, exception.NewNotFoundError("Data user belum ada di dalam sistem")
	}

	return res.FromUsers(users), nil
}

func (s *userServiceImpl) GetUserByID(id uint) (*res.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil || user.ID == 0 {
		return nil, exception.NewNotFoundError("User tidak ditemukan")
	}

	return res.FromUser(user), nil
}

func (s *userServiceImpl) CreateUser(r *req.UserCreateRequest) (*res.UserResponse, error) {
	newUser := &model.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		FullName: r.FullName,
		NIP:      r.NIP,
		Role:     r.Role,
	}

	if err := s.userRepo.Create(newUser); err != nil {
		return nil, err
	}

	return res.FromUser(newUser), nil
}
