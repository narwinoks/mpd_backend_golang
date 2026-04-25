package auth

import (
	"backend-app/internal/modules/auth/controller"
	tokenRepo "backend-app/internal/modules/auth/repository/personal_access_token"
	repo "backend-app/internal/modules/auth/repository/user"
	svcAuth "backend-app/internal/modules/auth/service/auth"
	svcUser "backend-app/internal/modules/auth/service/user"

	"github.com/google/wire"
)

var AuthProviderSet = wire.NewSet(
	repo.NewUserRepository,
	tokenRepo.NewTokenRepository,
	svcAuth.NewAuthService,
	svcUser.NewUserService,
	controller.NewAuthController,
	controller.NewUserController,
	NewAuthRouter,
)
