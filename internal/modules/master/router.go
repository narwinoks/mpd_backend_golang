package master

import (
	"backend-app/internal/modules/master/controller"
	"backend-app/internal/modules/master/controller/general"

	"github.com/gin-gonic/gin"
)

type MasterRouter struct {
	userController       *controller.UserController
	roleController       *controller.RoleController
	registryController   *controller.RegistryController
	permissionController *controller.PermissionController
	religionController   *general.ReligionController
	genderController     *general.GenderController
}

func NewMasterRouter(
	userController *controller.UserController,
	roleController *controller.RoleController,
	registryController *controller.RegistryController,
	permissionController *controller.PermissionController,
	religionController *general.ReligionController,
	genderController *general.GenderController,
) *MasterRouter {
	return &MasterRouter{
		userController:       userController,
		roleController:       roleController,
		registryController:   registryController,
		permissionController: permissionController,
		religionController:   religionController,
		genderController:     genderController,
	}
}

func (r *MasterRouter) RegisterRoutes(rg *gin.RouterGroup) {
	master := rg.Group("/master")
	{
		users := master.Group("/users")
		{
			users.GET("", r.userController.FindAll)
			users.GET("/:id", r.userController.FindByID)
			users.POST("", r.userController.Create)
		}

		roles := master.Group("/roles")
		{
			roles.GET("", r.roleController.FindAll)
			roles.GET("/:id", r.roleController.FindByID)
			roles.POST("", r.roleController.Create)
			roles.PUT("/:id", r.roleController.Update)
			roles.DELETE("/:id", r.roleController.Delete)
		}

		registries := master.Group("/registries")
		{
			registries.GET("/menu", r.registryController.GetMenu)
			registries.GET("", r.registryController.FindAll)
			registries.GET("/:id", r.registryController.FindByID)
			registries.POST("", r.registryController.Create)
			registries.PUT("/:id", r.registryController.Update)
			registries.DELETE("/:id", r.registryController.Delete)
		}

		permissions := master.Group("/permissions")
		{
			permissions.GET("", r.permissionController.FindAll)
			permissions.GET("/:id", r.permissionController.FindByID)
			permissions.POST("", r.permissionController.Create)
			permissions.PUT("/:id", r.permissionController.Update)
			permissions.DELETE("/:id", r.permissionController.Delete)
		}

		religions := master.Group("/general/religions")
		{
			religions.GET("", r.religionController.FindAll)
			religions.GET("/:id", r.religionController.FindByID)
			religions.POST("", r.religionController.Create)
			religions.PUT("/:id", r.religionController.Update)
			religions.DELETE("/:id", r.religionController.Delete)
		}

		genders := master.Group("/general/genders")
		{
			genders.GET("", r.genderController.FindAll)
			genders.GET("/:id", r.genderController.FindByID)
			genders.POST("", r.genderController.Create)
			genders.PUT("/:id", r.genderController.Update)
			genders.DELETE("/:id", r.genderController.Delete)
		}
	}
}
