package user

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/repository/user"
	res "backend-app/internal/modules/auth/response/user"

	"github.com/sirupsen/logrus"
)

type userServiceImpl struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}

func (s *userServiceImpl) GetProfile(userID uint32) (*res.ProfileResponse, error) {
	logrus.Infof("Fetching profile for user_id: %d", userID)

	user, err := s.repo.GetProfile(userID)
	if err != nil {
		logrus.Errorf("Failed to fetch user profile: %v", err)
		return nil, exception.NewNotFoundError("User not found")
	}

	// DDD: Mapping Domain Model to Response DTO
	response := &res.ProfileResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role: res.RoleInfo{
			ID:   user.Role.ID,
			Name: user.Role.Role,
		},
	}

	if user.Employee != nil {
		response.Employee = &res.EmployeeInfo{
			ID:             user.Employee.ID,
			FullName:       user.Employee.FullName,
			NIP:            user.Employee.NIP,
			IdentityNumber: user.Employee.IdentityNumber,
		}
	}

	return response, nil
}
