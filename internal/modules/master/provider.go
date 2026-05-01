package master

import (
	"backend-app/internal/modules/master/controller"
	repoReg "backend-app/internal/modules/master/repository/registry"
	repoRole "backend-app/internal/modules/master/repository/role"
	repoUser "backend-app/internal/modules/master/repository/user"
	repoPerm "backend-app/internal/modules/master/repository/permission"
	repoRel "backend-app/internal/modules/master/repository/general/religion"
	svcReg "backend-app/internal/modules/master/service/registry"
	svcRole "backend-app/internal/modules/master/service/role"
	svcUser "backend-app/internal/modules/master/service/user"
	svcPerm "backend-app/internal/modules/master/service/permission"
	svcRel "backend-app/internal/modules/master/service/general/religion"
	ctrlGen "backend-app/internal/modules/master/controller/general"

	"github.com/google/wire"
)

var MasterProviderSet = wire.NewSet(
	repoUser.NewUserRepository,
	repoRole.NewRoleRepository,
	repoReg.NewRegistryRepository,
	repoPerm.NewPermissionRepository,
	repoRel.NewReligionRepository,
	svcUser.NewUserService,
	svcRole.NewRoleService,
	svcReg.NewRegistryService,
	svcPerm.NewPermissionService,
	svcRel.NewReligionService,
	controller.NewUserController,
	controller.NewRoleController,
	controller.NewRegistryController,
	controller.NewPermissionController,
	ctrlGen.NewReligionController,
	NewMasterRouter,
)
