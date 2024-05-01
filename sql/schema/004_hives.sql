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

-- +goose Down
DROP TABLE hives;

