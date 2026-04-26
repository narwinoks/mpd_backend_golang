package auth

import (
	"backend-app/config"
	"backend-app/internal/modules/auth/controller"
	"backend-app/internal/modules/auth/middleware"
	"backend-app/internal/modules/auth/repository/module"
	"backend-app/internal/modules/auth/repository/permission"
	"backend-app/internal/modules/auth/repository/personal_access_token"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authController       *controller.AuthController
	userController       *controller.UserController
	menuController       *controller.MenuController
	permissionController *controller.PermissionController
	config               *config.Config
	tokenRepo            personal_access_token.TokenRepository
	permRepo             permission.PermissionRepository
	moduleRepo           module.ModuleRepository
}

func NewAuthRouter(
	authController *controller.AuthController,
	userController *controller.UserController,
	menuController *controller.MenuController,
	permissionController *controller.PermissionController,
	config *config.Config,
	tokenRepo personal_access_token.TokenRepository,
	permRepo permission.PermissionRepository,
	moduleRepo module.ModuleRepository,
) *AuthRouter {
	return &AuthRouter{
		authController:       authController,
		userController:       userController,
		menuController:       menuController,
		permissionController: permissionController,
		config:               config,
		tokenRepo:            tokenRepo,
		permRepo:             permRepo,
		moduleRepo:           moduleRepo,
	}
}

func (r *AuthRouter) RegisterRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", r.authController.Login)
		auth.POST("/refresh-token", r.authController.RefreshToken)
		auth.POST("/logout", middleware.AuthMiddleware(r.config, r.tokenRepo), r.authController.Logout)

		user := auth.Group("user")
		user.GET("/profile", middleware.AuthMiddleware(r.config, r.tokenRepo), r.userController.GetProfile)
		user.GET("/menus", middleware.AuthMiddleware(r.config, r.tokenRepo), r.menuController.GetSideMenu)
		user.GET("/permissions", middleware.AuthMiddleware(r.config, r.tokenRepo), r.permissionController.GetUserPermissions)

		// Example usage of Permission and Module middleware
		// user.GET("/example-permission", 
		//     middleware.AuthMiddleware(r.config, r.tokenRepo), 
		//     middleware.PermissionMiddleware(r.permRepo, "USER_VIEW"), 
		//     r.userController.GetProfile)
		
		// user.GET("/example-module", 
		//     middleware.AuthMiddleware(r.config, r.tokenRepo), 
		//     middleware.ModuleMiddleware(r.moduleRepo, "MOD_USER_MANAGEMENT"), 
		//     r.userController.GetProfile)
	}
}
