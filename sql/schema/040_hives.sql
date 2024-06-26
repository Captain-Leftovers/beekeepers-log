-- +goose Up
CREATE TABLE hives(
    id TEXT PRIMARY KEY,
    farm_id TEXT NOT NULL,
    name TEXT NOT NULL,
    isEmpty BOOLEAN NOT NULL DEFAULT 1,
    notes TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    last_inspection_date TIMESTAMP DEFAULT NULL,
    last_inspection_health TEXT DEFAULT NULL,
    CHECK (
        last_inspection_health IN ('healthy', 'weak', 'swarming', NULL)
    ) CHECK (isEmpty IN (0, 1)),
    FOREIGN KEY(farm_id) REFERENCES farms(id) ON DELETE CASCADE
);
CREATE INDEX farm_id_index ON hives(farm_id);
-- +goose Down
DROP TABLE hives;