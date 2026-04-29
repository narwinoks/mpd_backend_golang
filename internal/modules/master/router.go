package master

import (
	"backend-app/config"
	"backend-app/internal/modules/auth/middleware"
	"backend-app/internal/modules/auth/repository/personal_access_token"
	"backend-app/internal/modules/master/controller"

	"github.com/gin-gonic/gin"
)

type MasterRouter struct {
	userController *controller.UserController
	roleController *controller.RoleController
	config         *config.Config
	tokenRepo      personal_access_token.TokenRepository
}

func NewMasterRouter(
	userController *controller.UserController,
	roleController *controller.RoleController,
	config *config.Config,
	tokenRepo personal_access_token.TokenRepository,
) *MasterRouter {
	return &MasterRouter{
		userController: userController,
		roleController: roleController,
		config:         config,
		tokenRepo:      tokenRepo,
	}
}

func (r *MasterRouter) RegisterRoutes(rg *gin.RouterGroup) {
	master := rg.Group("/master", middleware.AuthMiddleware(r.config, r.tokenRepo))
	{
		users := master.Group("/users")
		{
			users.GET("/", r.userController.FindAll)
			users.GET("/:id", r.userController.FindByID)
			users.POST("/", r.userController.Create)
		}

		roles := master.Group("/roles")
		{
			roles.GET("/", r.roleController.FindAll)
			roles.GET("/:id", r.roleController.FindByID)
			roles.POST("/", r.roleController.Create)
			roles.PUT("/:id", r.roleController.Update)
			roles.DELETE("/:id", r.roleController.Delete)
		}
	}
}
