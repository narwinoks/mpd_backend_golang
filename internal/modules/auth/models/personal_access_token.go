package models

import (
	"backend-app/internal/base/models"
	"time"
)

type PersonalAccessToken struct {
	models.BaseModel
	UserID    uint32    `gorm:"column:user_id"`
	Token     string    `gorm:"column:token;type:text"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	IsRevoked bool      `gorm:"column:is_revoked;default:false"`
}

func (PersonalAccessToken) TableName() string {
	return "personal_access_tokens_m"
}
