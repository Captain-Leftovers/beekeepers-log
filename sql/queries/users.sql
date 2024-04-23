-- name: CreateUser :one
INSERT INTO users (id, username, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;
