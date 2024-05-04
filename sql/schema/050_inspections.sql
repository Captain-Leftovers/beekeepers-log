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
    FOREIGN KEY(hive_id) REFERENCES hives(id) ON DELETE CASCADE,
);

CREATE INDEX hive_id_index ON inspections(hive_id);
CREATE INDEX inspection_date_index ON inspections(inspection_date);


-- +goose Down
DROP TABLE inspections;


-- TODO see if you need to add farm_id to inspections table or not maybe the inspection id can be composite can  be composite