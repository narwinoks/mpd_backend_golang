-- 1. departments_m
CREATE TABLE IF NOT EXISTS departments_m (
    id          SERIAL PRIMARY KEY,
    uuid        CHAR(36)        NOT NULL,
    is_active   BOOLEAN         DEFAULT TRUE,
    profile_id  INTEGER         NOT NULL,
    external_code VARCHAR(15),
    created_at  TIMESTAMP(6),
    updated_at  TIMESTAMP(6),
    deleted_at  TIMESTAMP(6),
    created_by  INTEGER,
    updated_by  INTEGER,
    deleted_by  INTEGER,
    department_name VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS uidx_departments_uuid         ON departments_m(uuid);
CREATE INDEX IF NOT EXISTS idx_departments_profile_id    ON departments_m(profile_id);
CREATE INDEX IF NOT EXISTS idx_departments_is_active     ON departments_m(is_active);
CREATE INDEX IF NOT EXISTS idx_departments_external_code ON departments_m(external_code);

-- 2. wards_m
CREATE TABLE IF NOT EXISTS wards_m (
    id                   SERIAL PRIMARY KEY,
    uuid                 CHAR(36)     NOT NULL,
    is_active            BOOLEAN      DEFAULT TRUE,
    profile_id           INTEGER      NOT NULL,
    external_code        VARCHAR(20),
    created_at           TIMESTAMP(6),
    updated_at           TIMESTAMP(6),
    deleted_at           TIMESTAMP(6),
    created_by           INTEGER,
    updated_by           INTEGER,
    deleted_by           INTEGER,
    ward_name            VARCHAR(255) NOT NULL,
    department_id        INTEGER,
    is_executive         BOOLEAN      DEFAULT FALSE,
    icon                 VARCHAR(100),
    queue_number_prefix  VARCHAR(200),
    CONSTRAINT fk_wards_department FOREIGN KEY (department_id) REFERENCES departments_m(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS uidx_wards_uuid         ON wards_m(uuid);
CREATE INDEX IF NOT EXISTS idx_wards_profile_id    ON wards_m(profile_id);
CREATE INDEX IF NOT EXISTS idx_wards_is_active     ON wards_m(is_active);
CREATE INDEX IF NOT EXISTS idx_wards_external_code ON wards_m(external_code);
CREATE INDEX IF NOT EXISTS idx_wards_department_id ON wards_m(department_id);

-- 3. rooms_m
CREATE TABLE IF NOT EXISTS rooms_m (
    id                   SERIAL PRIMARY KEY,
    uuid                 CHAR(36)     NOT NULL,
    is_active            BOOLEAN      DEFAULT TRUE,
    profile_id           INTEGER      NOT NULL,
    external_code        VARCHAR(20),
    created_at           TIMESTAMP(6),
    updated_at           TIMESTAMP(6),
    deleted_at           TIMESTAMP(6),
    created_by           INTEGER,
    updated_by           INTEGER,
    deleted_by           INTEGER,
    ward_id              INTEGER,
    class_id             INTEGER,
    rs_online_code       VARCHAR(50),
    room_name            VARCHAR(255) NOT NULL,
    bed_count            INTEGER      NOT NULL DEFAULT 0,
    occupied_room_count  INTEGER      NOT NULL DEFAULT 0,
    available_room_count INTEGER      NOT NULL DEFAULT 0,
    CONSTRAINT fk_rooms_ward FOREIGN KEY (ward_id) REFERENCES wards_m(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS uidx_rooms_uuid         ON rooms_m(uuid);
CREATE INDEX IF NOT EXISTS idx_rooms_profile_id    ON rooms_m(profile_id);
CREATE INDEX IF NOT EXISTS idx_rooms_is_active     ON rooms_m(is_active);
CREATE INDEX IF NOT EXISTS idx_rooms_external_code ON rooms_m(external_code);
CREATE INDEX IF NOT EXISTS idx_rooms_ward_id       ON rooms_m(ward_id);
CREATE INDEX IF NOT EXISTS idx_rooms_class_id      ON rooms_m(class_id);

-- 4. beds_m
CREATE TABLE IF NOT EXISTS beds_m (
    id             SERIAL PRIMARY KEY,
    uuid           CHAR(36)     NOT NULL,
    is_active      BOOLEAN      DEFAULT TRUE,
    profile_id     INTEGER      NOT NULL,
    external_code  VARCHAR(20),
    created_at     TIMESTAMP(6),
    updated_at     TIMESTAMP(6),
    deleted_at     TIMESTAMP(6),
    created_by     INTEGER,
    updated_by     INTEGER,
    deleted_by     INTEGER,
    room_id        INTEGER,
    bed_status_id  INTEGER,
    bed_number     VARCHAR(255) NOT NULL,
    description    TEXT         NOT NULL DEFAULT '',
    merged_bed_id  INTEGER,
    CONSTRAINT fk_beds_room FOREIGN KEY (room_id) REFERENCES rooms_m(id)
);

CREATE UNIQUE INDEX IF NOT EXISTS uidx_beds_uuid         ON beds_m(uuid);
CREATE INDEX IF NOT EXISTS idx_beds_profile_id    ON beds_m(profile_id);
CREATE INDEX IF NOT EXISTS idx_beds_is_active     ON beds_m(is_active);
CREATE INDEX IF NOT EXISTS idx_beds_external_code ON beds_m(external_code);
CREATE INDEX IF NOT EXISTS idx_beds_room_id       ON beds_m(room_id);
CREATE INDEX IF NOT EXISTS idx_beds_bed_status_id ON beds_m(bed_status_id);
