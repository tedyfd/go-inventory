-- name: CreateProduct :one
INSERT INTO product (id, created_at, updated_at, name, quantity, user_id, seller_id, category_id) 
VALUES ($1, $2, $3, $4, $5,$6, $7, $8)
RETURNING *;

-- name: GetProduct :many
SELECT product.id, 
	product.name, 
	product.quantity,
	product.created_at, 
	product.updated_at, 
	category.name as category_name,
	seller.name as seller_name,
	users.name as user_name
	FROM product
INNER JOIN category ON product.category_id = category.id
INNER JOIN seller ON product.seller_id = seller.id
INNER JOIN users ON product.user_id = users.id;

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

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;