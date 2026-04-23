package auth

import (
	"backend-app/internal/modules/auth/controller"
	repo "backend-app/internal/modules/auth/repository/user"
	svc "backend-app/internal/modules/auth/service/user"

	"github.com/google/wire"
)

var AuthProviderSet = wire.NewSet(
	repo.NewUserRepository,
	svc.NewUserService,
	controller.NewAuthController,
	NewAuthRouter,
)
