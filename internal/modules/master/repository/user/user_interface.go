package user

import (
	"backend-app/internal/modules/auth/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByUsername(username string) (bool, error)
	FindByEmail(email string) (bool, error)
	FindByNIP(nip string) (bool, error)
	Create(user *models.User) error
}
