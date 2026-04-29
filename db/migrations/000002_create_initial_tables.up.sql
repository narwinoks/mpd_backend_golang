-- Phase 1: Table Creation
CREATE TABLE religions_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    religion VARCHAR(100)
);

CREATE TABLE profiles_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    province_id INTEGER,
    city_id INTEGER,
    subdistrict_id INTEGER,
    village_id INTEGER,
    postal_code VARCHAR(10),
    email VARCHAR(100),
    name VARCHAR(100),
    profile TEXT,
    government_name VARCHAR(100),
    phone VARCHAR(20),
    telp VARCHAR(20),
    full_address TEXT
);


CREATE TABLE provinces_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    province VARCHAR(100)
);

CREATE TABLE cities_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    province_id INTEGER, city VARCHAR(100)
);

CREATE TABLE subdistrict_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    city_id INTEGER, province_id INTEGER, subdistrict VARCHAR(100)
);

CREATE TABLE genders_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    gender VARCHAR(20)
);

CREATE TABLE job_categories_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    code VARCHAR(20), job_category VARCHAR(100)
);

CREATE TABLE job_titles_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    job_category_id INTEGER, code VARCHAR(20), job_title VARCHAR(100)
);

CREATE TABLE employment_statuses_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    code VARCHAR(20), employee_status VARCHAR(100)
);

CREATE TABLE marital_status_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    material_status VARCHAR(50)
);

CREATE TABLE employee_group_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    employee_group VARCHAR(100)
);

CREATE TABLE positions_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    position VARCHAR(100)
);

CREATE TABLE employees_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    religion_id INTEGER,
    gender_id INTEGER,
    job_title_id INTEGER,
    employment_status_id INTEGER,
    full_name VARCHAR(100),
    identity_number VARCHAR(20),
    nip VARCHAR(20),
    npwp VARCHAR(20),
    birth_place VARCHAR(100),
    birth_date DATE
);

CREATE TABLE employee_details_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    employee_id INTEGER UNIQUE,
    marital_status_id INTEGER,
    functional_position_id INTEGER,
    structural_position_id INTEGER,
    join_date DATE,
    resign_date DATE,
    retirement_date DATE
);

CREATE TABLE roles_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    role VARCHAR(100)
);

CREATE TABLE users_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    role_id INTEGER,
    employee_id INTEGER,
    username VARCHAR(50),
    email VARCHAR(100),
    password VARCHAR(255)
);

CREATE TABLE app_modules_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    code VARCHAR(20), name VARCHAR(100), category VARCHAR(50), sort_order INTEGER
);

CREATE TABLE app_menus_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    app_module_id INTEGER,
    parent_id INTEGER,
    code VARCHAR(100),
    name VARCHAR(100),
    path VARCHAR(255),
    description TEXT,
    icon VARCHAR(100),
    sort_order INTEGER
);

CREATE TABLE role_modules_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    role_id INTEGER, modules_id INTEGER
);

CREATE TABLE user_modules_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    user_id INTEGER, modules_id INTEGER
);

CREATE TABLE app_permission_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    permission VARCHAR(100)
);

CREATE TABLE role_permission_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    role_id INTEGER, permission_id INTEGER
);

CREATE TABLE user_permission_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    user_id INTEGER, permission_id INTEGER
);

-- Phase 2: Alter Constraints
ALTER TABLE religions_m ADD CONSTRAINT fk_religions_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE religions_m ADD CONSTRAINT fk_religions_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE religions_m ADD CONSTRAINT fk_religions_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE profiles_m ADD CONSTRAINT fk_profiles_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE profiles_m ADD CONSTRAINT fk_profiles_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE profiles_m ADD CONSTRAINT fk_profiles_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE provinces_m ADD CONSTRAINT fk_provinces_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE provinces_m ADD CONSTRAINT fk_provinces_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE provinces_m ADD CONSTRAINT fk_provinces_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE cities_m ADD CONSTRAINT fk_cities_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE cities_m ADD CONSTRAINT fk_cities_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE cities_m ADD CONSTRAINT fk_cities_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE subdistrict_m ADD CONSTRAINT fk_subdistrict_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE subdistrict_m ADD CONSTRAINT fk_subdistrict_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE subdistrict_m ADD CONSTRAINT fk_subdistrict_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE genders_m ADD CONSTRAINT fk_genders_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE genders_m ADD CONSTRAINT fk_genders_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE genders_m ADD CONSTRAINT fk_genders_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE job_categories_m ADD CONSTRAINT fk_job_categories_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE job_categories_m ADD CONSTRAINT fk_job_categories_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE job_categories_m ADD CONSTRAINT fk_job_categories_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE job_titles_m ADD CONSTRAINT fk_job_titles_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE job_titles_m ADD CONSTRAINT fk_job_titles_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE job_titles_m ADD CONSTRAINT fk_job_titles_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employment_statuses_m ADD CONSTRAINT fk_employment_statuses_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employment_statuses_m ADD CONSTRAINT fk_employment_statuses_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employment_statuses_m ADD CONSTRAINT fk_employment_statuses_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE marital_status_m ADD CONSTRAINT fk_marital_status_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE marital_status_m ADD CONSTRAINT fk_marital_status_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE marital_status_m ADD CONSTRAINT fk_marital_status_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_group_m ADD CONSTRAINT fk_employee_group_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_group_m ADD CONSTRAINT fk_employee_group_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_group_m ADD CONSTRAINT fk_employee_group_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE positions_m ADD CONSTRAINT fk_positions_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE positions_m ADD CONSTRAINT fk_positions_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE positions_m ADD CONSTRAINT fk_positions_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employees_m ADD CONSTRAINT fk_employees_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employees_m ADD CONSTRAINT fk_employees_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employees_m ADD CONSTRAINT fk_employees_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE roles_m ADD CONSTRAINT fk_roles_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE roles_m ADD CONSTRAINT fk_roles_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE roles_m ADD CONSTRAINT fk_roles_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE users_m ADD CONSTRAINT fk_users_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE users_m ADD CONSTRAINT fk_users_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE users_m ADD CONSTRAINT fk_users_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_modules_m ADD CONSTRAINT fk_app_modules_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_modules_m ADD CONSTRAINT fk_app_modules_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_modules_m ADD CONSTRAINT fk_app_modules_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_menus_m ADD CONSTRAINT fk_app_menus_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_menus_m ADD CONSTRAINT fk_app_menus_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_menus_m ADD CONSTRAINT fk_app_menus_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE role_modules_m ADD CONSTRAINT fk_role_modules_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE role_modules_m ADD CONSTRAINT fk_role_modules_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE role_modules_m ADD CONSTRAINT fk_role_modules_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE user_modules_m ADD CONSTRAINT fk_user_modules_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE user_modules_m ADD CONSTRAINT fk_user_modules_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE user_modules_m ADD CONSTRAINT fk_user_modules_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_permission_m ADD CONSTRAINT fk_app_permission_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_permission_m ADD CONSTRAINT fk_app_permission_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_permission_m ADD CONSTRAINT fk_app_permission_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE role_permission_m ADD CONSTRAINT fk_role_permission_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE role_permission_m ADD CONSTRAINT fk_role_permission_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE role_permission_m ADD CONSTRAINT fk_role_permission_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE user_permission_m ADD CONSTRAINT fk_user_permission_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE user_permission_m ADD CONSTRAINT fk_user_permission_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE user_permission_m ADD CONSTRAINT fk_user_permission_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE cities_m ADD CONSTRAINT fk_cities_m_province_id FOREIGN KEY (province_id) REFERENCES provinces_m(id) ON DELETE SET NULL;
ALTER TABLE subdistrict_m ADD CONSTRAINT fk_subdistrict_m_city_id FOREIGN KEY (city_id) REFERENCES cities_m(id) ON DELETE SET NULL;
ALTER TABLE subdistrict_m ADD CONSTRAINT fk_subdistrict_m_province_id FOREIGN KEY (province_id) REFERENCES provinces_m(id) ON DELETE SET NULL;
ALTER TABLE employees_m ADD CONSTRAINT fk_employees_m_religion_id FOREIGN KEY (religion_id) REFERENCES religions_m(id) ON DELETE SET NULL;
ALTER TABLE employees_m ADD CONSTRAINT fk_employees_m_gender_id FOREIGN KEY (gender_id) REFERENCES genders_m(id) ON DELETE SET NULL;
ALTER TABLE employees_m ADD CONSTRAINT fk_employees_m_employment_status_id FOREIGN KEY (employment_status_id) REFERENCES employment_statuses_m(id) ON DELETE SET NULL;
ALTER TABLE employees_m ADD CONSTRAINT fk_employees_m_job_title_id FOREIGN KEY (job_title_id) REFERENCES job_titles_m(id) ON DELETE SET NULL;
ALTER TABLE job_titles_m ADD CONSTRAINT fk_job_titles_m_job_category_id FOREIGN KEY (job_category_id) REFERENCES job_categories_m(id) ON DELETE SET NULL;
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_employee_id FOREIGN KEY (employee_id) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_marital_status_id FOREIGN KEY (marital_status_id) REFERENCES marital_status_m(id) ON DELETE SET NULL;
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_functional_position_id FOREIGN KEY (functional_position_id) REFERENCES positions_m(id) ON DELETE SET NULL;
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_structural_position_id FOREIGN KEY (structural_position_id) REFERENCES positions_m(id) ON DELETE SET NULL;
ALTER TABLE users_m ADD CONSTRAINT fk_users_m_role_id FOREIGN KEY (role_id) REFERENCES roles_m(id) ON DELETE SET NULL;
ALTER TABLE users_m ADD CONSTRAINT fk_users_m_employee_id FOREIGN KEY (employee_id) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE app_menus_m ADD CONSTRAINT fk_app_menus_m_app_module_id FOREIGN KEY (app_module_id) REFERENCES app_modules_m(id) ON DELETE SET NULL;
ALTER TABLE app_menus_m ADD CONSTRAINT fk_app_menus_m_parent_id FOREIGN KEY (parent_id) REFERENCES app_menus_m(id) ON DELETE SET NULL;
ALTER TABLE role_modules_m ADD CONSTRAINT fk_role_modules_m_role_id FOREIGN KEY (role_id) REFERENCES roles_m(id) ON DELETE SET NULL;
ALTER TABLE role_modules_m ADD CONSTRAINT fk_role_modules_m_modules_id FOREIGN KEY (modules_id) REFERENCES app_modules_m(id) ON DELETE SET NULL;
ALTER TABLE user_modules_m ADD CONSTRAINT fk_user_modules_m_user_id FOREIGN KEY (user_id) REFERENCES users_m(id) ON DELETE SET NULL;
ALTER TABLE user_modules_m ADD CONSTRAINT fk_user_modules_m_modules_id FOREIGN KEY (modules_id) REFERENCES app_modules_m(id) ON DELETE SET NULL;
ALTER TABLE role_permission_m ADD CONSTRAINT fk_role_permission_m_role_id FOREIGN KEY (role_id) REFERENCES roles_m(id) ON DELETE SET NULL;
ALTER TABLE role_permission_m ADD CONSTRAINT fk_role_permission_m_permission_id FOREIGN KEY (permission_id) REFERENCES app_permission_m(id) ON DELETE SET NULL;
ALTER TABLE user_permission_m ADD CONSTRAINT fk_user_permission_m_user_id FOREIGN KEY (user_id) REFERENCES users_m(id) ON DELETE SET NULL;
ALTER TABLE user_permission_m ADD CONSTRAINT fk_user_permission_m_permission_id FOREIGN KEY (permission_id) REFERENCES app_permission_m(id) ON DELETE SET NULL;

-- Profile ID References
ALTER TABLE religions_m ADD CONSTRAINT fk_religions_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE profiles_m ADD CONSTRAINT fk_profiles_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE provinces_m ADD CONSTRAINT fk_provinces_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE cities_m ADD CONSTRAINT fk_cities_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE subdistrict_m ADD CONSTRAINT fk_subdistrict_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE genders_m ADD CONSTRAINT fk_genders_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE job_categories_m ADD CONSTRAINT fk_job_categories_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE job_titles_m ADD CONSTRAINT fk_job_titles_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE employment_statuses_m ADD CONSTRAINT fk_employment_statuses_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE marital_status_m ADD CONSTRAINT fk_marital_status_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE employee_group_m ADD CONSTRAINT fk_employee_group_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE positions_m ADD CONSTRAINT fk_positions_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE employees_m ADD CONSTRAINT fk_employees_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE roles_m ADD CONSTRAINT fk_roles_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE users_m ADD CONSTRAINT fk_users_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE app_modules_m ADD CONSTRAINT fk_app_modules_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE app_menus_m ADD CONSTRAINT fk_app_menus_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE role_modules_m ADD CONSTRAINT fk_role_modules_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE user_modules_m ADD CONSTRAINT fk_user_modules_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE app_permission_m ADD CONSTRAINT fk_app_permission_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE role_permission_m ADD CONSTRAINT fk_role_permission_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE user_permission_m ADD CONSTRAINT fk_user_permission_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
