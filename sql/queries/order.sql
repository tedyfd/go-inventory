-- name: CreateOrder :one
INSERT INTO "order"(
    id, created_at, updated_at, customer_id
    ) 
VALUES ($1, $2, $3, $4)
RETURNING *;
