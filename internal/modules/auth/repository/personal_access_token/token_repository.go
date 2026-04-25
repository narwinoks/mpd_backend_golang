package personal_access_token

import (
	"backend-app/internal/modules/auth/models"
	"gorm.io/gorm"
)

type tokenRepositoryImpl struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepositoryImpl{db: db}
}

func (r *tokenRepositoryImpl) Create(token *models.PersonalAccessToken) error {
	return r.db.Create(token).Error
}

func (r *tokenRepositoryImpl) FindByToken(token string) (*models.PersonalAccessToken, error) {
	var m models.PersonalAccessToken
	err := r.db.Where("token = ?", token).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *tokenRepositoryImpl) RevokeByToken(token string) error {
	return r.db.Model(&models.PersonalAccessToken{}).
		Where("token = ?", token).
		Update("is_revoked", true).Error
}

func (r *tokenRepositoryImpl) IsRevoked(token string) (bool, error) {
	var m models.PersonalAccessToken
	err := r.db.Where("token = ?", token).First(&m).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return m.IsRevoked, nil
}
