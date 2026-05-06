CREATE TABLE IF NOT EXISTS funding_sources_m (
    id          SERIAL PRIMARY KEY,
    uuid        CHAR(36)        NOT NULL,
    is_active   BOOLEAN         DEFAULT TRUE,
    profile_id  INTEGER         NOT NULL,
    external_code VARCHAR(20),
    created_at  TIMESTAMP(6),
    updated_at  TIMESTAMP(6),
    deleted_at  TIMESTAMP(6),
    created_by  INTEGER,
    updated_by  INTEGER,
    deleted_by  INTEGER,
    funding_source VARCHAR(100) NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS uidx_funding_sources_uuid         ON funding_sources_m(uuid);
CREATE INDEX IF NOT EXISTS idx_funding_sources_profile_id    ON funding_sources_m(profile_id);
CREATE INDEX IF NOT EXISTS idx_funding_sources_is_active     ON funding_sources_m(is_active);
CREATE INDEX IF NOT EXISTS idx_funding_sources_external_code ON funding_sources_m(external_code);
