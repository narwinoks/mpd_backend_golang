CREATE TABLE master_registries_m (
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
    name VARCHAR(100) NOT NULL,
    path VARCHAR(100),
    icon VARCHAR(50),
    head_id INTEGER,
    sort_order INTEGER DEFAULT 0
);

CREATE INDEX idx_master_registries_uuid ON master_registries_m(uuid);
CREATE INDEX idx_master_registries_head_id ON master_registries_m(head_id);
