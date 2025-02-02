-- name: CreateOrderDetail :one
INSERT INTO order_detail(
    created_at, updated_at, 
    order_id, quantity, product_id
    ) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateOrderDetailAndUpdateProduct :one
SELECT create_order_detail_and_update_quantity($1, $2, $3);