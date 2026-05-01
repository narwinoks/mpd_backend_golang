package user

import (
	"backend-app/internal/modules/auth/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := r.db.WithContext(ctx).Preload("Role", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "uuid", "role")
	}).Preload("Employee", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "uuid", "full_name")
	}).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
func (r *userRepositoryImpl) FindByUsername(ctx context.Context, username string) (bool, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("username = ?", username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (bool, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (r *userRepositoryImpl) FindByNIP(ctx context.Context, nip string) (bool, error) {
	if nip == "" {
		return false, nil
	}
	var user models.User
	result := r.db.WithContext(ctx).Where("nip = ?", nip).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}
