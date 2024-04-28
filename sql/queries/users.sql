-- name: CreateUser :one
INSERT INTO users (id, username, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;


-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = ?
LIMIT 1;


-- name: GetUserById :one
SELECT * FROM users 
WHERE id = ?
LIMIT 1;


-- name: UpdateUser :one
UPDATE users
SET username = ?, email = ?, password = ?, updated_at = ?
WHERE id = ?
RETURNING *;