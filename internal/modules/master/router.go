package master

import (
	"backend-app/internal/modules/master/controller"

	"github.com/gin-gonic/gin"
)

type MasterRouter struct {
	userController *controller.UserController
}

func NewMasterRouter(userController *controller.UserController) *MasterRouter {
	return &MasterRouter{userController: userController}
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
	}
}
