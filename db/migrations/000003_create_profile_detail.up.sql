CREATE TABLE profile_detail_m (
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
     website varchar(100),
     longitude float,
     latitude float,
     registration_date date,
     moto TEXT
);

ALTER TABLE profile_detail_m ADD CONSTRAINT fk_profile_detail_m_created_by FOREIGN KEY (created_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE profile_detail_m ADD CONSTRAINT fk_profile_detail_m_updated_by FOREIGN KEY (updated_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE profile_detail_m ADD CONSTRAINT fk_profile_detail_m_deleted_by FOREIGN KEY (deleted_by) REFERENCES employees_m(id) ON DELETE SET NULL;
ALTER TABLE profile_detail_m ADD CONSTRAINT fk_profile_detail_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
