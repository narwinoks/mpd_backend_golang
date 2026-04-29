package auth

import (
	"backend-app/internal/modules/auth/controller"
	"backend-app/internal/modules/auth/middleware"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authController       *controller.AuthController
	userController       *controller.UserController
	menuController       *controller.MenuController
	permissionController *controller.PermissionController
	AuthMiddleware       *middleware.AuthMiddleware
	moduleMiddleware     *middleware.ModuleMiddleware
	permMiddleware       *middleware.PermissionMiddleware
}

func NewAuthRouter(
	authController *controller.AuthController,
	userController *controller.UserController,
	menuController *controller.MenuController,
	permissionController *controller.PermissionController,
	authMiddleware *middleware.AuthMiddleware,
	moduleMiddleware *middleware.ModuleMiddleware,
	permMiddleware *middleware.PermissionMiddleware,
) *AuthRouter {
	return &AuthRouter{
		authController:       authController,
		userController:       userController,
		menuController:       menuController,
		permissionController: permissionController,
		AuthMiddleware:       authMiddleware,
		moduleMiddleware:     moduleMiddleware,
		permMiddleware:       permMiddleware,
	}
}

func (r *AuthRouter) RegisterRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", r.authController.Login)
		auth.POST("/refresh-token", r.authController.RefreshToken)
		auth.POST("/logout", r.AuthMiddleware.Handle(), r.authController.Logout)

		user := auth.Group("user")
		user.Use(r.AuthMiddleware.Handle())
		{
			user.GET("/profile", r.userController.GetProfile)
			user.GET("/menus", r.menuController.GetSideMenu)
			user.GET("/permissions", r.permissionController.GetUserPermissions)
		}
	}
}
