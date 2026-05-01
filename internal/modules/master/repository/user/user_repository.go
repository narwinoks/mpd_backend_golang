package user

import (
	"backend-app/internal/modules/auth/models"
	"backend-app/pkg/pagination"
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

type UserWithCount struct {
	models.User
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *userRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]models.User, int64, error) {
	var results []UserWithCount
	var users []models.User
	var total int64

	err := r.db.WithContext(ctx).Model(&models.User{}).
		Preload("Role", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "role")
		}).
		Preload("Employee", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "uuid", "full_name")
		}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("username", "email")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			users = append(users, res.User)
		}
	}

	return users, total, nil
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
