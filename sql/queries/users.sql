-- name: CreateUser :one
INSERT INTO users (
    id, created_at, updated_at, username, password, name, api_key
    ) 
VALUES (
    $1, $2, $3, $4, $5, $6, 
    encode(sha256(random()::text::bytea), 'hex')
    )
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;