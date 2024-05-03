-- +goose Up
CREATE TABLE farms(
    id TEXT PRIMARY KEY,
    name TEXT  NOT NULL,   
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE farms;
