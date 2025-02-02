-- name: CreateSeller :one
INSERT INTO seller (
    id, created_at, updated_at, name
    ) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSeller :many
SELECT * FROM seller;

-- name: GetSellerByName :one
SELECT * FROM seller WHERE name=$1;

-- name: UpdateSeller :one
UPDATE seller SET
name = $1,
updated_at = $2
WHERE id = $3
RETURNING *;

-- name: DeleteSeller :exec
DELETE FROM seller WHERE id=$1;
