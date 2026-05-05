package master

import (
	"backend-app/internal/modules/master/controller"
	"backend-app/internal/modules/master/controller/app"
	"backend-app/internal/modules/master/controller/department"
	"backend-app/internal/modules/master/controller/employee"
	"backend-app/internal/modules/master/controller/general"
	"backend-app/internal/modules/master/controller/location"

	"github.com/gin-gonic/gin"
)

type MasterRouter struct {
	userController             *controller.UserController
	roleController             *controller.RoleController
	registryController         *controller.RegistryController
	appModuleController        *app.AppModuleController
	appMenuController          *app.AppMenuController
	permissionController       *controller.PermissionController
	religionController         *general.ReligionController
	genderController           *general.GenderController
	educationController        *general.EducationController
	bankController             *general.BankController
	maritalStatusController    *general.MaritalStatusController
	employeeController         *employee.EmployeeController
	provinceController         *location.ProvinceController
	cityController             *location.CityController
	subdistrictController      *location.SubdistrictController
	villageController          *location.VillageController
	jobCategoryController      *employee.JobCategoryController
	jobTitleController         *employee.JobTitleController
	positionController         *employee.PositionController
	employmentStatusController *employee.EmploymentStatusController
	departmentController       *department.DepartmentController
	wardController             *department.WardController
	roomController             *department.RoomController
	bedController              *department.BedController
}

func NewMasterRouter(
	userController *controller.UserController,
	roleController *controller.RoleController,
	registryController *controller.RegistryController,
	appModuleController *app.AppModuleController,
	permissionController *controller.PermissionController,
	religionController *general.ReligionController,
	genderController *general.GenderController,
	educationController *general.EducationController,
	bankController *general.BankController,
	maritalStatusController *general.MaritalStatusController,
	employeeController *employee.EmployeeController,
	provinceController *location.ProvinceController,
	cityController *location.CityController,
	subdistrictController *location.SubdistrictController,
	villageController *location.VillageController,
	jobCategoryController *employee.JobCategoryController,
	jobTitleController *employee.JobTitleController,
	positionController *employee.PositionController,
	employmentStatusController *employee.EmploymentStatusController,
	departmentController *department.DepartmentController,
	wardController *department.WardController,
	roomController *department.RoomController,
	bedController *department.BedController,
) *MasterRouter {
	return &MasterRouter{
		userController:             userController,
		roleController:             roleController,
		registryController:         registryController,
		appModuleController:        appModuleController,
		permissionController:       permissionController,
		religionController:         religionController,
		genderController:           genderController,
		educationController:        educationController,
		bankController:             bankController,
		maritalStatusController:    maritalStatusController,
		employeeController:         employeeController,
		provinceController:         provinceController,
		cityController:             cityController,
		subdistrictController:      subdistrictController,
		villageController:          villageController,
		jobCategoryController:      jobCategoryController,
		jobTitleController:         jobTitleController,
		positionController:         positionController,
		employmentStatusController: employmentStatusController,
		departmentController:       departmentController,
		wardController:             wardController,
		roomController:             roomController,
		bedController:              bedController,
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

		appModules := master.Group("/app/app-modules")
		{
			appModules.GET("", r.appModuleController.FindAll)
			appModules.GET("/:id", r.appModuleController.FindByID)
			appModules.POST("", r.appModuleController.Create)
			appModules.PUT("/:id", r.appModuleController.Update)
			appModules.DELETE("/:id", r.appModuleController.Delete)
		}

		appMenus := master.Group("/app/app-menus")
		{
			appMenus.GET("", r.appMenuController.FindAll)
			appMenus.GET("/:id", r.appMenuController.FindByID)
			appMenus.POST("", r.appMenuController.Create)
			appMenus.PUT("/:id", r.appMenuController.Update)
			appMenus.DELETE("/:id", r.appMenuController.Delete)
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

		educations := master.Group("/general/educations")
		{
			educations.GET("", r.educationController.FindAll)
			educations.GET("/:id", r.educationController.FindByID)
			educations.POST("", r.educationController.Create)
			educations.PUT("/:id", r.educationController.Update)
			educations.DELETE("/:id", r.educationController.Delete)
		}

		banks := master.Group("/general/bank")
		{
			banks.GET("", r.bankController.FindAll)
			banks.GET("/:id", r.bankController.FindByID)
			banks.POST("", r.bankController.Create)
			banks.PUT("/:id", r.bankController.Update)
			banks.DELETE("/:id", r.bankController.Delete)
		}

		employees := master.Group("/employee")
		{
			employees.GET("", r.employeeController.FindAll)
			employees.GET("/:id", r.employeeController.FindByID)
			employees.POST("", r.employeeController.Create)
			employees.PUT("/:id", r.employeeController.Update)
			employees.DELETE("/:id", r.employeeController.Delete)
		}

		maritalStatuses := master.Group("/general/marital-status")
		{
			maritalStatuses.GET("", r.maritalStatusController.FindAll)
			maritalStatuses.GET("/:id", r.maritalStatusController.FindByID)
			maritalStatuses.POST("", r.maritalStatusController.Create)
			maritalStatuses.PUT("/:id", r.maritalStatusController.Update)
			maritalStatuses.DELETE("/:id", r.maritalStatusController.Delete)
		}

		provinces := master.Group("/location/provinces")
		{
			provinces.GET("", r.provinceController.FindAll)
			provinces.GET("/:id", r.provinceController.FindByID)
			provinces.POST("", r.provinceController.Create)
			provinces.PUT("/:id", r.provinceController.Update)
			provinces.DELETE("/:id", r.provinceController.Delete)
		}

		cities := master.Group("/location/cities")
		{
			cities.GET("", r.cityController.FindAll)
			cities.GET("/:id", r.cityController.FindByID)
			cities.POST("", r.cityController.Create)
			cities.PUT("/:id", r.cityController.Update)
			cities.DELETE("/:id", r.cityController.Delete)
		}

		subdistricts := master.Group("/location/subdistricts")
		{
			subdistricts.GET("", r.subdistrictController.FindAll)
			subdistricts.GET("/:id", r.subdistrictController.FindByID)
			subdistricts.POST("", r.subdistrictController.Create)
			subdistricts.PUT("/:id", r.subdistrictController.Update)
			subdistricts.DELETE("/:id", r.subdistrictController.Delete)
		}

		villages := master.Group("/location/villages")
		{
			villages.GET("", r.villageController.FindAll)
			villages.GET("/:id", r.villageController.FindByID)
			villages.POST("", r.villageController.Create)
			villages.PUT("/:id", r.villageController.Update)
			villages.DELETE("/:id", r.villageController.Delete)
		}

		jobCategories := master.Group("/employee/job-categories")
		{
			jobCategories.GET("", r.jobCategoryController.FindAll)
			jobCategories.GET("/:id", r.jobCategoryController.FindByID)
			jobCategories.POST("", r.jobCategoryController.Create)
			jobCategories.PUT("/:id", r.jobCategoryController.Update)
			jobCategories.DELETE("/:id", r.jobCategoryController.Delete)
		}

		jobTitles := master.Group("/employee/job-titles")
		{
			jobTitles.GET("", r.jobTitleController.FindAll)
			jobTitles.GET("/:id", r.jobTitleController.FindByID)
			jobTitles.POST("", r.jobTitleController.Create)
			jobTitles.PUT("/:id", r.jobTitleController.Update)
			jobTitles.DELETE("/:id", r.jobTitleController.Delete)
		}

		positions := master.Group("/employee/positions")
		{
			positions.GET("", r.positionController.FindAll)
			positions.GET("/:id", r.positionController.FindByID)
			positions.POST("", r.positionController.Create)
			positions.PUT("/:id", r.positionController.Update)
			positions.DELETE("/:id", r.positionController.Delete)
		}

		employmentStatuses := master.Group("/employee/employment-statuses")
		{
			employmentStatuses.GET("", r.employmentStatusController.FindAll)
			employmentStatuses.GET("/:id", r.employmentStatusController.FindByID)
			employmentStatuses.POST("", r.employmentStatusController.Create)
			employmentStatuses.PUT("/:id", r.employmentStatusController.Update)
			employmentStatuses.DELETE("/:id", r.employmentStatusController.Delete)
		}

		departments := master.Group("/department/departments")
		{
			departments.GET("", r.departmentController.FindAll)
			departments.GET("/:id", r.departmentController.FindByID)
			departments.POST("", r.departmentController.Create)
			departments.PUT("/:id", r.departmentController.Update)
			departments.DELETE("/:id", r.departmentController.Delete)
		}

		wards := master.Group("/department/wards")
		{
			wards.GET("", r.wardController.FindAll)
			wards.GET("/:id", r.wardController.FindByID)
			wards.POST("", r.wardController.Create)
			wards.PUT("/:id", r.wardController.Update)
			wards.DELETE("/:id", r.wardController.Delete)
		}

		rooms := master.Group("/department/rooms")
		{
			rooms.GET("", r.roomController.FindAll)
			rooms.GET("/:id", r.roomController.FindByID)
			rooms.POST("", r.roomController.Create)
			rooms.PUT("/:id", r.roomController.Update)
			rooms.DELETE("/:id", r.roomController.Delete)
		}

		beds := master.Group("/department/beds")
		{
			beds.GET("", r.bedController.FindAll)
			beds.GET("/:id", r.bedController.FindByID)
			beds.POST("", r.bedController.Create)
			beds.PUT("/:id", r.bedController.Update)
			beds.DELETE("/:id", r.bedController.Delete)
		}
	}
}
