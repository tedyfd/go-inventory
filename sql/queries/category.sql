-- name: CreateCategory :one
INSERT INTO category (id, name, description) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1;