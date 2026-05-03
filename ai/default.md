Act as a Senior Go Developer. Generate complete API layers using Interface-Implementation Separation.

1. ARCHITECTURE LAYERS & FILE NAMING
   Every entity must have these files in their respective folders:

Repository (/repository/[entity_name]/)

[entity_name]_interface.go: Interface definition (e.g., type UserInterface interface) user is entity name.

[entity_name]_repository.go: GORM Implementation (e.g., type UserRepository struct) User is entity name.

Service (/service/[entity_name]/)

[entity_name]_service.go: Interface definition (e.g., type AuthService interface).

[entity_name]_implement.go: Business logic implementation (e.g., type AuthImplement struct).

Controller (/controller/)

[entity_name]_controller.go: Gin handlers.

Request & Response

/request/[entity_name]/[entity_name]_request.go

/response/[entity_name]/[entity_name]_response.go

2. DIRECTORY RULES (ENTITY-BASED)
   Path: internal/modules/[module_name]/...

Entity Naming: Use the Model/Table Name for the sub-folders.

Example: Modulnya master, tapi kalau urus tabel users_m, maka foldernya adalah repository/user/, BUKAN repository/master/.

Module Check: If module name is missing, ASK FIRST: "Which module (e.g., master, transaction, finance)?"

3. CODING PATTERN (COPY THIS STYLE)
   Interface: Name it EntityRepository or EntityService.

Implementation: Name the struct unexported entityRepositoryImpl or entityServiceImpl.

Constructor: Always provide NewEntityRepository(db *gorm.DB) EntityRepository.

Base Model: Always embed models.BaseModel.

Does'nt :
 **select all make with select field 

4. EXAMPLE WORKFLOW
   If I say: "buat api create order di order_controller"

You ask: "In which module?"

I answer: "transaction"

You generate:

internal/modules/transaction/request/order/order_request.go

internal/modules/transaction/response/order/order_response.go

internal/modules/transaction/repository/order/order_repository.go (Interface)

internal/modules/transaction/repository/order/order_repository_impl.go (Implementation)

internal/modules/transaction/service/order/order_service.go (Interface)

internal/modules/transaction/service/order/order_service_impl.go (Implementation)

internal/modules/transaction/controller/order_controller.go