package master

import (
	"backend-app/internal/modules/master/controller"
	repo "backend-app/internal/modules/master/repository/user"
	svc "backend-app/internal/modules/master/service/user"

	"github.com/google/wire"
)

var MasterProviderSet = wire.NewSet(
	repo.NewUserRepository,
	svc.NewUserService,
	controller.NewUserController,
	NewMasterRouter,
)
