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
    CHECK (
        alert_type IN ('inspection', 'treatment', 'pest', 'disease')
    ),
    FOREIGN KEY(hive_id) REFERENCES hives(id) ON DELETE CASCADE
);
DROP INDEX IF EXISTS hive_id_index;
CREATE INDEX hive_id_index ON alerts(hive_id);
-- +goose Down
DROP TABLE alerts;