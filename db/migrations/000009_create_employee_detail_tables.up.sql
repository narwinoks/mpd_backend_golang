-- 1. banks_m
CREATE TABLE banks_m (
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
    bank VARCHAR(20)
);

-- 2. add field employee_details_m
ALTER TABLE employee_details_m ADD COLUMN bank_id INTEGER;
ALTER TABLE employee_details_m ADD COLUMN bank_account_number VARCHAR(50);
ALTER TABLE employee_details_m ADD COLUMN bank_account_name VARCHAR(100);

-- 3. educations_m
CREATE TABLE educations_m (
    id SERIAL PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    profile_id INTEGER NOT NULL,
    external_code VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    education_type VARCHAR(20) NOT NULL CHECK (education_type IN ('FORMAL', 'INFORMAL', 'OTHER')),
    code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    sort_order INT DEFAULT 0
);

-- 4. employee_addresses_m
CREATE TABLE employee_addresses_m (
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
    employee_id INTEGER NOT NULL,
    address_type VARCHAR(50) NOT NULL ,
    full_address TEXT NOT NULL ,
    province_id INTEGER,
    city_id INTEGER,
    subdistrict_id INTEGER,
    village_id INTEGER
);

-- 5. employee_educations_m
CREATE TABLE employee_educations_m (
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
    employee_id INTEGER NOT NULL,
    education_level_id INTEGER NOT NULL,
    institution_name VARCHAR(200),
    institution_address VARCHAR(250),
    major VARCHAR(255),
    start_date DATE,
    graduation_date DATE,
    certificate_date DATE,
    certificate_number VARCHAR(255),
    gpa DECIMAL(10,2),
    front_title VARCHAR(30),
    back_title VARCHAR(30),
    is_highest BOOLEAN DEFAULT FALSE
);

-- Foreign Key Constraints

-- banks_m FKs
ALTER TABLE banks_m ADD CONSTRAINT fk_banks_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE banks_m ADD CONSTRAINT fk_banks_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE banks_m ADD CONSTRAINT fk_banks_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE banks_m ADD CONSTRAINT fk_banks_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;

-- employee_details_m bank FK
ALTER TABLE employee_details_m ADD CONSTRAINT fk_employee_details_m_bank_id FOREIGN KEY (bank_id) REFERENCES banks_m(id) ON DELETE SET NULL;

-- educations_m FKs
ALTER TABLE educations_m ADD CONSTRAINT fk_educations_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE educations_m ADD CONSTRAINT fk_educations_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE educations_m ADD CONSTRAINT fk_educations_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE educations_m ADD CONSTRAINT fk_educations_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;

-- employee_addresses_m FKs
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_employee_id FOREIGN KEY (employee_id) REFERENCES employees_m(id) ON DELETE CASCADE;
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_province_id FOREIGN KEY (province_id) REFERENCES provinces_m(id) ON DELETE SET NULL;
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_city_id FOREIGN KEY (city_id) REFERENCES cities_m(id) ON DELETE SET NULL;
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_subdistrict_id FOREIGN KEY (subdistrict_id) REFERENCES subdistrict_m(id) ON DELETE SET NULL;
ALTER TABLE employee_addresses_m ADD CONSTRAINT fk_employee_addresses_m_village_id FOREIGN KEY (village_id) REFERENCES villages_m(id) ON DELETE SET NULL;

-- employee_educations_m FKs
ALTER TABLE employee_educations_m ADD CONSTRAINT fk_employee_educations_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
ALTER TABLE employee_educations_m ADD CONSTRAINT fk_employee_educations_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_educations_m ADD CONSTRAINT fk_employee_educations_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_educations_m ADD CONSTRAINT fk_employee_educations_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE employee_educations_m ADD CONSTRAINT fk_employee_educations_m_employee_id FOREIGN KEY (employee_id) REFERENCES employees_m(id) ON DELETE CASCADE;
ALTER TABLE employee_educations_m ADD CONSTRAINT fk_employee_educations_m_education_level_id FOREIGN KEY (education_level_id) REFERENCES educations_m(id) ON DELETE SET NULL;
