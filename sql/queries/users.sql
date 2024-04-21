-- name: CreateUser :one
INSERT INTO users (id, username, email, created_at, updated_at)
VALUES (?, ?, ?, NOW(), NOW())
RETURNING *;
