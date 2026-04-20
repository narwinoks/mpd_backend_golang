package user

import (
	"backend-app/internal/modules/master/model"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(id uint) (*model.User, error)
	Create(user *model.User) error
}
