-- +goose Up
CREATE TABLE test_table(
    id text PRIMARY KEY,
    description TEXT DEFAULT ''
);
-- +goose Down
DROP TABLE test_table;