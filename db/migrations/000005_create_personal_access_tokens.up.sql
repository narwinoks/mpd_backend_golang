CREATE TABLE personal_access_tokens_m (
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
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    is_revoked BOOLEAN DEFAULT FALSE
);

ALTER TABLE personal_access_tokens_m ADD CONSTRAINT fk_personal_access_tokens_m_user_id FOREIGN KEY (user_id) REFERENCES users_m(id) ON DELETE CASCADE;
ALTER TABLE personal_access_tokens_m ADD CONSTRAINT fk_personal_access_tokens_m_profile_id FOREIGN KEY (profile_id) REFERENCES profiles_m(id) ON DELETE SET NULL;
