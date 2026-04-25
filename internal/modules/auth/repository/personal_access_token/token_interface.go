package personal_access_token

import "backend-app/internal/modules/auth/models"

type TokenRepository interface {
	Create(token *models.PersonalAccessToken) error
	FindByToken(token string) (*models.PersonalAccessToken, error)
	RevokeByToken(token string) error
	IsRevoked(token string) (bool, error)
}
