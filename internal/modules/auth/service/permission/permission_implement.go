package permission

import (
	repo "backend-app/internal/modules/auth/repository/permission"
	"backend-app/internal/modules/auth/repository/user"

	"github.com/sirupsen/logrus"
)

type permissionServiceImpl struct {
	repo     repo.PermissionRepository
	userRepo user.UserRepository
}

func NewPermissionService(repo repo.PermissionRepository, userRepo user.UserRepository) PermissionService {
	return &permissionServiceImpl{repo: repo, userRepo: userRepo}
}

func (s *permissionServiceImpl) GetUserPermissions(userID uint32) ([]string, error) {
	// We need role_id
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	logrus.Infof("Fetching permissions for user_id: %d, role_id: %d", userID, user.RoleID)

	permissions, err := s.repo.GetUserPermissions(userID, user.RoleID)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
