package master

import (
	"backend-app/internal/modules/auth/repository/personal_access_token"
	"backend-app/internal/modules/master/controller"
	repoRole "backend-app/internal/modules/master/repository/role"
	repoUser "backend-app/internal/modules/master/repository/user"
	svcRole "backend-app/internal/modules/master/service/role"
	svcUser "backend-app/internal/modules/master/service/user"

	"github.com/google/wire"
)

var MasterProviderSet = wire.NewSet(
	repoUser.NewUserRepository,
	repoRole.NewRoleRepository,
	svcUser.NewUserService,
	svcRole.NewRoleService,
	controller.NewUserController,
	controller.NewRoleController,
	personal_access_token.NewTokenRepository,
	NewMasterRouter,
)
