package master

import (
	"backend-app/internal/modules/master/controller"
	ctrlEmp "backend-app/internal/modules/master/controller/employee"
	ctrlGen "backend-app/internal/modules/master/controller/general"
	repoEmpStatus "backend-app/internal/modules/master/repository/employee/employment_status"
	repoJobCat "backend-app/internal/modules/master/repository/employee/job_category"
	repoJobTitle "backend-app/internal/modules/master/repository/employee/job_title"
	repoGender "backend-app/internal/modules/master/repository/general/gender"
	repoRel "backend-app/internal/modules/master/repository/general/religion"
	repoPerm "backend-app/internal/modules/master/repository/permission"
	repoReg "backend-app/internal/modules/master/repository/registry"
	repoRole "backend-app/internal/modules/master/repository/role"
	repoUser "backend-app/internal/modules/master/repository/user"
	svcEmpStatus "backend-app/internal/modules/master/service/employee/employment_status"
	svcJobCat "backend-app/internal/modules/master/service/employee/job_category"
	svcJobTitle "backend-app/internal/modules/master/service/employee/job_title"
	svcGender "backend-app/internal/modules/master/service/general/gender"
	svcRel "backend-app/internal/modules/master/service/general/religion"
	svcPerm "backend-app/internal/modules/master/service/permission"
	svcReg "backend-app/internal/modules/master/service/registry"
	svcRole "backend-app/internal/modules/master/service/role"
	svcUser "backend-app/internal/modules/master/service/user"

	"github.com/google/wire"
)

var MasterProviderSet = wire.NewSet(
	repoUser.NewUserRepository,
	repoRole.NewRoleRepository,
	repoReg.NewRegistryRepository,
	repoPerm.NewPermissionRepository,
	repoRel.NewReligionRepository,
	repoGender.NewGenderRepository,
	repoJobCat.NewJobCategoryRepository,
	repoJobTitle.NewJobTitleRepository,
	repoEmpStatus.NewEmploymentStatusRepository,
	svcUser.NewUserService,
	svcRole.NewRoleService,
	svcReg.NewRegistryService,
	svcPerm.NewPermissionService,
	svcRel.NewReligionService,
	svcGender.NewGenderService,
	svcJobCat.NewJobCategoryService,
	svcJobTitle.NewJobTitleService,
	svcEmpStatus.NewEmploymentStatusService,
	controller.NewUserController,
	controller.NewRoleController,
	controller.NewRegistryController,
	controller.NewPermissionController,
	ctrlGen.NewReligionController,
	ctrlGen.NewGenderController,
	ctrlEmp.NewJobCategoryController,
	ctrlEmp.NewJobTitleController,
	ctrlEmp.NewEmploymentStatusController,
	NewMasterRouter,
)
