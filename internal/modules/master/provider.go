package master

import (
	"backend-app/internal/modules/master/controller"
	ctrlDept "backend-app/internal/modules/master/controller/department"
	ctrlEmp "backend-app/internal/modules/master/controller/employee"
	ctrlGen "backend-app/internal/modules/master/controller/general"
	ctrlLoc "backend-app/internal/modules/master/controller/location"
	repoBed "backend-app/internal/modules/master/repository/department/bed"
	repoDept "backend-app/internal/modules/master/repository/department/department"
	repoRoom "backend-app/internal/modules/master/repository/department/room"
	repoWard "backend-app/internal/modules/master/repository/department/ward"
	repoEmployee "backend-app/internal/modules/master/repository/employee"
	repoEmpStatus "backend-app/internal/modules/master/repository/employee/employment_status"
	repoJobCat "backend-app/internal/modules/master/repository/employee/job_category"
	repoJobTitle "backend-app/internal/modules/master/repository/employee/job_title"
	repoPosition "backend-app/internal/modules/master/repository/employee/position"
	repoBank "backend-app/internal/modules/master/repository/general/bank"
	repoEducation "backend-app/internal/modules/master/repository/general/education"
	repoGender "backend-app/internal/modules/master/repository/general/gender"
	repoMaritalStatus "backend-app/internal/modules/master/repository/general/marital_status"
	repoRel "backend-app/internal/modules/master/repository/general/religion"
	repoProvince "backend-app/internal/modules/master/repository/location/province"
	repoCity "backend-app/internal/modules/master/repository/location/city"
	repoPerm "backend-app/internal/modules/master/repository/permission"
	repoReg "backend-app/internal/modules/master/repository/registry"
	repoRole "backend-app/internal/modules/master/repository/role"
	repoUser "backend-app/internal/modules/master/repository/user"
	svcBed "backend-app/internal/modules/master/service/department/bed"
	svcDept "backend-app/internal/modules/master/service/department/department"
	svcRoom "backend-app/internal/modules/master/service/department/room"
	svcWard "backend-app/internal/modules/master/service/department/ward"
	svcEmployee "backend-app/internal/modules/master/service/employee"
	svcEmpStatus "backend-app/internal/modules/master/service/employee/employment_status"
	svcJobCat "backend-app/internal/modules/master/service/employee/job_category"
	svcJobTitle "backend-app/internal/modules/master/service/employee/job_title"
	svcPosition "backend-app/internal/modules/master/service/employee/position"
	svcBank "backend-app/internal/modules/master/service/general/bank"
	svcEducation "backend-app/internal/modules/master/service/general/education"
	svcGender "backend-app/internal/modules/master/service/general/gender"
	svcMaritalStatus "backend-app/internal/modules/master/service/general/marital_status"
	svcRel "backend-app/internal/modules/master/service/general/religion"
	svcProvince "backend-app/internal/modules/master/service/location/province"
	svcCity "backend-app/internal/modules/master/service/location/city"
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
	repoEducation.NewEducationRepository,
	repoBank.NewBankRepository,
	repoMaritalStatus.NewMaritalStatusRepository,
	repoJobCat.NewJobCategoryRepository,
	repoJobTitle.NewJobTitleRepository,
	repoPosition.NewPositionRepository,
	repoEmployee.NewEmployeeRepository,
	repoEmpStatus.NewEmploymentStatusRepository,
	repoDept.NewDepartmentRepository,
	repoWard.NewWardRepository,
	repoRoom.NewRoomRepository,
	repoBed.NewBedRepository,
	repoProvince.NewProvinceRepository,
	repoCity.NewCityRepository,
	svcUser.NewUserService,
	svcRole.NewRoleService,
	svcReg.NewRegistryService,
	svcPerm.NewPermissionService,
	svcRel.NewReligionService,
	svcGender.NewGenderService,
	svcEducation.NewEducationService,
	svcBank.NewBankService,
	svcMaritalStatus.NewMaritalStatusService,
	svcJobCat.NewJobCategoryService,
	svcJobTitle.NewJobTitleService,
	svcPosition.NewPositionService,
	svcEmployee.NewEmployeeService,
	svcEmpStatus.NewEmploymentStatusService,
	svcDept.NewDepartmentService,
	svcWard.NewWardService,
	svcRoom.NewRoomService,
	svcBed.NewBedService,
	svcProvince.NewProvinceService,
	svcCity.NewCityService,
	controller.NewUserController,
	controller.NewRoleController,
	controller.NewRegistryController,
	controller.NewPermissionController,
	ctrlGen.NewReligionController,
	ctrlGen.NewGenderController,
	ctrlGen.NewEducationController,
	ctrlGen.NewBankController,
	ctrlGen.NewMaritalStatusController,
	ctrlEmp.NewJobCategoryController,
	ctrlEmp.NewJobTitleController,
	ctrlEmp.NewPositionController,
	ctrlEmp.NewEmployeeController,
	ctrlEmp.NewEmploymentStatusController,
	ctrlDept.NewDepartmentController,
	ctrlDept.NewWardController,
	ctrlDept.NewRoomController,
	ctrlDept.NewBedController,
	ctrlLoc.NewProvinceController,
	ctrlLoc.NewCityController,
	NewMasterRouter,
)
