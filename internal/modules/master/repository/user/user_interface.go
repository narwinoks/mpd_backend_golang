package user

import (
	"backend-app/internal/modules/auth/models"
	"backend-app/pkg/pagination"
	"context"
)

type UserRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]models.User, int64, error)
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (bool, error)
	FindByEmail(ctx context.Context, email string) (bool, error)
	FindByNIP(ctx context.Context, nip string) (bool, error)
	Create(ctx context.Context, user *models.User) error
}
