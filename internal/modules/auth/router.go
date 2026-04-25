package auth

import (
	"backend-app/config"
	"backend-app/internal/modules/auth/controller"
	"backend-app/internal/modules/auth/middleware"
	"backend-app/internal/modules/auth/repository/personal_access_token"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authController *controller.AuthController
	userController *controller.UserController
	config         *config.Config
	tokenRepo      personal_access_token.TokenRepository
}

func NewAuthRouter(
	authController *controller.AuthController,
	userController *controller.UserController,
	config *config.Config,
	tokenRepo personal_access_token.TokenRepository,
) *AuthRouter {
	return &AuthRouter{
		authController: authController,
		userController: userController,
		config:         config,
		tokenRepo:      tokenRepo,
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
	}
}
