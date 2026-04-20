package user

import (
	"backend-app/internal/modules/master/model/user"
)

type UserRepository interface {
	FindAll() ([]user.User, error)
	FindByID(id uint) (*user.User, error)
	FindByUsername(username string) (bool, error)
	FindByEmail(email string) (bool, error)
	FindByNIP(nip string) (bool, error)
	Create(user *user.User) error
}
