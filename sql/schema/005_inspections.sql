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

-- +goose Down
DROP TABLE inspections;