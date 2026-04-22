Act as a Senior Backend Engineer (Go & PostgreSQL Expert). I need you to generate two things:

Golang Structs using GORM based on the Mermaid ERD provided.

A single, comprehensive Raw SQL Migration File (.sql) for PostgreSQL.

DO NOT provide AutoMigrate code. The goal is a manual, rock-solid migration strategy.

1. GORM STRUCT RULES
   Base Model: Create base_model.go with a BaseModel struct to be embedded in all others.

Fields: ID (uint32), UUID (char 36), IsActive (bool), ProfileID (*uint32), ExternalCode (varchar 20), CreatedAt, UpdatedAt, DeletedAt, and Audit pointers: CreatedBy (*uint32), UpdatedBy (*uint32), DeletedBy (*uint32).

Naming: Use PascalCase for Structs. Use gorm:"column:..." tags (snake_case). Tables must end in _m.

Pointer Strategy: All Foreign Keys referring to EmployeesM (Audit fields or direct links) MUST be pointers (*uint32) to allow nullability and avoid circular reference issues in the code.

Table Names: Implement the TableName() string method for every struct to ensure the _m suffix is preserved.

2. SQL MIGRATION RULES (THE 'SET NULL' STRATEGY)
   Generate a sequential SQL script that follows this logical flow:

Phase 1: Table Creation. Create all 23 tables.

CRITICAL: Define created_by, updated_by, and deleted_by as INTEGER NULL (no constraints yet).

Initialize other FKs (like province_id) as nullable as well to ensure the script runs from start to finish without dependency errors.

Phase 2: Alter Constraints. After all tables are created, use ALTER TABLE ... ADD CONSTRAINT ... FOREIGN KEY ... REFERENCES for all relationships.

For created_by, updated_by, and deleted_by, set them to ON DELETE SET NULL.

3. FILE ORGANIZATION
   Organize the output into these specific paths:

internal/base/models/base_model.go

internal/auth/models/ (For: Roles, Users, Permissions, Modules, Menus, Pivot tables)

internal/master/models/employee/ (For: Employees, Details, Religions, Genders, Titles, etc.)

db/migrations/up and down .sql

erDiagram
%% ==================================================
%% RELATIONSHIPS
%% ==================================================

    PROVINCES_M ||--o{ CITIES_M : "province_id"
    CITIES_M ||--o{ SUBDISTRICT_M : "city_id"
    PROVINCES_M ||--o{ SUBDISTRICT_M : "province_id"
    
    RELIGIONS_M ||--o{ EMPLOYEES_M : "religion_id"
    GENDERS_M ||--o{ EMPLOYEES_M : "gender_id"
    EMPLOYMENT_STATUSES_M ||--o{ EMPLOYEES_M : "employment_status_id"
    JOB_TITLES_M ||--o{ EMPLOYEES_M : "job_title_id"
    JOB_CATEGORIES_M ||--o{ JOB_TITLES_M : "job_category_id"

    EMPLOYEES_M ||--|| EMPLOYEE_DETAILS_M : "employee_id"
    MARITAL_STATUS_M ||--o{ EMPLOYEE_DETAILS_M : "marital_status_id"
    POSITIONS_M ||--o{ EMPLOYEE_DETAILS_M : "functional_position_id"
    POSITIONS_M ||--o{ EMPLOYEE_DETAILS_M : "structural_position_id"

    ROLES_M ||--o{ USERS_M : "role_id"
    EMPLOYEES_M ||--o{ USERS_M : "employee_id"

    APP_MODULES_M ||--o{ APP_MENUS_M : "app_module_id"
    APP_MENUS_M ||--o{ APP_MENUS_M : "parent_id"

    ROLES_M ||--o{ ROLE_MODULES_M : "role_id"
    APP_MODULES_M ||--o{ ROLE_MODULES_M : "modules_id"

    USERS_M ||--o{ USER_MODULES_M : "user_id"
    APP_MODULES_M ||--o{ USER_MODULES_M : "modules_id"

    APP_PERMISSION_M ||--o{ ROLE_PERMISSION_M : "permission_id"
    ROLES_M ||--o{ ROLE_PERMISSION_M : "role_id"

    APP_PERMISSION_M ||--o{ USER_PERMISSION_M : "permission_id"
    USERS_M ||--o{ USER_PERMISSION_M : "user_id"

    %% ==================================================
    %% GLOBAL TEMPLATE DEFINITIONS (Applied to all)
    %% ==================================================

    RELIGIONS_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string religion
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    PROFILES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int province_id
        int city_id
        int subdistrict_id
        int village_id
        string postal_code
        string email
        string name
        string profile
        string government_name
        string phone
        string telp
        text full_address
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    PROVINCES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string province
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    CITIES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int province_id
        string city
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    SUBDISTRICT_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int city_id
        int province_id
        string subdistrict
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    GENDERS_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string gender
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    JOB_CATEGORIES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string code
        string job_category
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    JOB_TITLES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int job_category_id
        string code
        string job_title
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    EMPLOYMENT_STATUSES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string code
        string employee_status
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    MARITAL_STATUS_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string material_status
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    EMPLOYEE_GROUP_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string employee_group
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    POSITIONS_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string position
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    EMPLOYEES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int religion_id
        int gender_id
        int job_title_id
        int employment_status_id
        string full_name
        string identity_number
        string nip
        string npwp
        string birth_place
        date birth_date
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    EMPLOYEE_DETAILS_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int employee_id
        int marital_status_id
        int functional_position_id
        int structural_position_id
        date join_date
        date resign_date
        date retirement_date
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    ROLES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string role
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    USERS_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int role_id
        int employee_id
        string username
        string email
        string password
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    APP_MODULES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string code
        string name
        string category
        int sort_order
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    APP_MENUS_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int app_module_id
        int parent_id
        string code
        string name
        string path
        string description
        string icon
        int sort_order
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    ROLE_MODULES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int role_id
        int modules_id
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    USER_MODULES_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int user_id
        int modules_id
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    APP_PERMISSION_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        string permission
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    ROLE_PERMISSION_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int role_id
        int permission_id
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }

    USER_PERMISSION_M {
        int id
        string uuid
        bool is_active
        int profile_id
        string external_code
        int user_id
        int permission_id
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
        int created_by
        int updated_by
        int deleted_by
    }
