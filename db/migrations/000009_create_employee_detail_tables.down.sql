ALTER TABLE employee_educations_m DROP CONSTRAINT IF EXISTS fk_employee_educations_m_education_level_id;
ALTER TABLE employee_educations_m DROP CONSTRAINT IF EXISTS fk_employee_educations_m_employee_id;
ALTER TABLE employee_educations_m DROP CONSTRAINT IF EXISTS fk_employee_educations_m_deleted_by;
ALTER TABLE employee_educations_m DROP CONSTRAINT IF EXISTS fk_employee_educations_m_updated_by;
ALTER TABLE employee_educations_m DROP CONSTRAINT IF EXISTS fk_employee_educations_m_created_by;
ALTER TABLE employee_educations_m DROP CONSTRAINT IF EXISTS fk_employee_educations_m_profile_id;
DROP TABLE IF EXISTS employee_educations_m;

ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_village_id;
ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_subdistrict_id;
ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_city_id;
ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_province_id;
ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_employee_id;
ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_deleted_by;
ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_updated_by;
ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_created_by;
ALTER TABLE employee_addresses_m DROP CONSTRAINT IF EXISTS fk_employee_addresses_m_profile_id;
DROP TABLE IF EXISTS employee_addresses_m;

ALTER TABLE educations_m DROP CONSTRAINT IF EXISTS fk_educations_m_deleted_by;
ALTER TABLE educations_m DROP CONSTRAINT IF EXISTS fk_educations_m_updated_by;
ALTER TABLE educations_m DROP CONSTRAINT IF EXISTS fk_educations_m_created_by;
ALTER TABLE educations_m DROP CONSTRAINT IF EXISTS fk_educations_m_profile_id;
DROP TABLE IF EXISTS educations_m;

ALTER TABLE employee_details_m DROP CONSTRAINT IF EXISTS fk_employee_details_m_bank_id;
ALTER TABLE employee_details_m DROP COLUMN IF EXISTS bank_account_name;
ALTER TABLE employee_details_m DROP COLUMN IF EXISTS bank_account_number;
ALTER TABLE employee_details_m DROP COLUMN IF EXISTS bank_id;

ALTER TABLE banks_m DROP CONSTRAINT IF EXISTS fk_banks_m_deleted_by;
ALTER TABLE banks_m DROP CONSTRAINT IF EXISTS fk_banks_m_updated_by;
ALTER TABLE banks_m DROP CONSTRAINT IF EXISTS fk_banks_m_created_by;
ALTER TABLE banks_m DROP CONSTRAINT IF EXISTS fk_banks_m_profile_id;
DROP TABLE IF EXISTS banks_m;
