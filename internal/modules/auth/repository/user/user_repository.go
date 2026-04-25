package user

import (
	"backend-app/internal/modules/auth/models"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindByIdentity(identity string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ? OR email = ?", identity, identity).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindByID(id uint32) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetProfile(userID uint32) (*models.User, error) {
	var user models.User
	err := r.db.
		Select("id", "username", "email", "role_id", "employee_id").
		Preload("Employee", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "full_name", "identity_number", "nip")
		}).
		Preload("Role", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "role")
		}).First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
