package auth

import (
	"backend-app/internal/modules/auth/controller"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authController *controller.AuthController
}

func NewAuthRouter(authController *controller.AuthController) *AuthRouter {
	return &AuthRouter{authController: authController}
}

func (r *AuthRouter) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", r.authController.Login)
	}
}
