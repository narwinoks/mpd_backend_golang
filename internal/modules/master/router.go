package master

import (
	"backend-app/internal/modules/master/controller"

	"github.com/gin-gonic/gin"
)

type MasterRouter struct {
	userController *controller.UserController
	roleController *controller.RoleController
}

func NewMasterRouter(
	userController *controller.UserController,
	roleController *controller.RoleController,
) *MasterRouter {
	return &MasterRouter{
		userController: userController,
		roleController: roleController,
	}
}

func (r *MasterRouter) RegisterRoutes(rg *gin.RouterGroup) {
	master := rg.Group("/master")
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
