//go:build wireinject
// +build wireinject

package auth

import (
	"backend-app/config"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeAuthRouter(cfg *config.Config, db *gorm.DB) *AuthRouter {
	wire.Build(
		AuthProviderSet,
	)
	return &AuthRouter{}
}
