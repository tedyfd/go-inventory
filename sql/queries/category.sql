-- name: CreateCategory :one
INSERT INTO category (id, name, description) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1;

-- name: GetCategory :many
SELECT * FROM category;

-- name: GetCategoryByName :one
SELECT * FROM category WHERE name=$1;

-- name: UpdateCategory :one
UPDATE category SET
name = $1,
description = $2
WHERE id = $3
RETURNING *;
