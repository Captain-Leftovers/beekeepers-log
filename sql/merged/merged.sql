-- +goose Up
CREATE TABLE users(
    id TEXT PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT CHECK(role IN ('user', 'admin')) DEFAULT 'user' NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;
-- +goose Up
CREATE TABLE farms(
    id TEXT PRIMARY KEY,
    name TEXT  NOT NULL,   
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE farms;
-- +goose Up
CREATE TABLE user_farms(
    user_id TEXT NOT NULL,
    farm_id TEXT NOT NULL,
    PRIMARY KEY(user_id, farm_id),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(farm_id) REFERENCES farms(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE user_farms;

-- +goose Up
CREATE TABLE hives(
    id TEXT PRIMARY KEY,
    farm_id TEXT NOT NULL,
    name TEXT NOT NULL,
    isEmpty BOOLEAN NOT NULL DEFAULT 1,
    notes TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CHECK (isEmpty IN (0, 1)),
    FOREIGN KEY(farm_id) REFERENCES farms(id) ON DELETE CASCADE
);

CREATE INDEX farm_id_index ON hives(farm_id);


-- +goose Down
DROP TABLE hives;

-- +goose Up
CREATE TABLE inspections(
    id text PRIMARY KEY,
    hive_id TEXT NOT NULL,
    inspection_date TIMESTAMP NOT NULL,
    honey_stores TEXT NOT NULL,
    pests TEXT NOT NULL DEFAULT 'none',
    diseases TEXT NOT NULL DEFAULT 'none',
    queen_seen BOOLEAN NOT NULL DEFAULT 0,
    eggs_present BOOLEAN NOT NULL DEFAULT 0,
    larvae_present BOOLEAN NOT NULL DEFAULT 0,
    capped_brood BOOLEAN NOT NULL DEFAULT 0,
    notes TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    overall_health TEXT NOT NULL DEFAULT 'healthy',
    CHECK (honey_stores IN ('low', 'medium', 'high')),
    CHECK (pests IN ('none', 'few', 'many')),
    CHECK (diseases IN ('none', 'few', 'many')),
    CHECK (overall_health IN ('healthy', 'weak','swarming')),
    CHECK (queen_seen IN (0, 1)),
    CHECK (eggs_present IN (0, 1)),
    CHECK (larvae_present IN (0, 1)),
    CHECK (capped_brood IN (0, 1)),
    FOREIGN KEY(hive_id) REFERENCES hives(id) ON DELETE CASCADE
);

CREATE INDEX hive_id_index ON inspections(hive_id);
CREATE INDEX inspection_date_index ON inspections(inspection_date);

-- +goose Down
DROP TABLE inspections;

-- +goose Up
CREATE TABLE alerts(
    id TEXT PRIMARY KEY,
    hive_id TEXT NOT NULL,
    alert_time TIMESTAMP NOT NULL,
    alert_type TEXT NOT NULL DEFAULT 'inspection',
    notes TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CHECK (status IN ('pending', 'closed')),
    CHECK (alert_type IN ('inspection', 'treatment', 'pest', 'disease')),
    FOREIGN KEY(hive_id) REFERENCES hives(id) ON DELETE CASCADE
);
CREATE INDEX hive_id_index ON alerts(hive_id);


-- +goose Down
DROP TABLE alerts;




-- below is seed data

