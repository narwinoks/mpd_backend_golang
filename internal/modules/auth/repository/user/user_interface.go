package user

import (
	"backend-app/internal/modules/auth/models"
)

type UserRepository interface {
	FindByIdentity(identity string) (*models.User, error)
	FindByID(id uint32) (*models.User, error)
}
