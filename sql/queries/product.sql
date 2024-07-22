-- name: CreateProduct :one
INSERT INTO product (id, created_at, updated_at, name, quantity, user_id, category_id) 
VALUES ($1, $2, $3, $4, $5,$6, $7)
RETURNING *;

-- name: GetProduct :many
SELECT * FROM product
INNER JOIN category ON product.category_id = category.id;

-- name: GetProductByID :one
SELECT * FROM product
INNER JOIN category ON product.category_id = category.id
WHERE product.id = $1;

-- name: UpdateProduct :one
UPDATE product 
SET updated_at = NOW(), 
name = $1,
quantity = $2
WHERE id = $3
RETURNING *;

-- NAME: DeleteProduct :exec
DELETE FROM product WHERE id = $1;