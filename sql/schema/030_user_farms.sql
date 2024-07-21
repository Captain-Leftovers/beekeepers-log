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