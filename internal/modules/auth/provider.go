package auth

import (
	"backend-app/internal/modules/auth/controller"
	repoMenu "backend-app/internal/modules/auth/repository/menu"
	tokenRepo "backend-app/internal/modules/auth/repository/personal_access_token"
	repoUser "backend-app/internal/modules/auth/repository/user"
	svcAuth "backend-app/internal/modules/auth/service/auth"
	svcMenu "backend-app/internal/modules/auth/service/menu"
	svcUser "backend-app/internal/modules/auth/service/user"

	"github.com/google/wire"
)

var AuthProviderSet = wire.NewSet(
	repoUser.NewUserRepository,
	tokenRepo.NewTokenRepository,
	repoMenu.NewMenuRepository,
	svcAuth.NewAuthService,
	svcUser.NewUserService,
	svcMenu.NewMenuService,
	controller.NewAuthController,
	controller.NewUserController,
	controller.NewMenuController,
	NewAuthRouter,
)
