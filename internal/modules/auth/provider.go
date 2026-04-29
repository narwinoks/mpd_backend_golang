package auth

import (
	"backend-app/internal/modules/auth/controller"
	"backend-app/internal/modules/auth/middleware"
	repoMenu "backend-app/internal/modules/auth/repository/menu"
	repoModule "backend-app/internal/modules/auth/repository/module"
	repoPerm "backend-app/internal/modules/auth/repository/permission"
	tokenRepo "backend-app/internal/modules/auth/repository/personal_access_token"
	repoUser "backend-app/internal/modules/auth/repository/user"
	svcAuth "backend-app/internal/modules/auth/service/auth"
	svcMenu "backend-app/internal/modules/auth/service/menu"
	svcPerm "backend-app/internal/modules/auth/service/permission"
	svcUser "backend-app/internal/modules/auth/service/user"

	"github.com/google/wire"
)

var AuthProviderSet = wire.NewSet(
	repoUser.NewUserRepository,
	tokenRepo.NewTokenRepository,
	repoMenu.NewMenuRepository,
	repoModule.NewModuleRepository,
	repoPerm.NewPermissionRepository,
	svcAuth.NewAuthService,
	svcUser.NewUserService,
	svcMenu.NewMenuService,
	svcPerm.NewPermissionService,
	controller.NewAuthController,
	controller.NewUserController,
	controller.NewMenuController,
	controller.NewPermissionController,
	middleware.NewAuthMiddleware,
	middleware.NewModuleMiddleware,
	middleware.NewPermissionMiddleware,
	NewAuthRouter,
)
