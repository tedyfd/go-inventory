-- name: CreateCustomer :one
INSERT INTO customer(
    id, created_at, updated_at, name
    ) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCustomer :many
SELECT * FROM customer;

-- name: GetCustomerByName :one
SELECT * FROM customer WHERE name = $1;

-- name: GetCustomerByID :one
SELECT * FROM customer WHERE id = $1;

-- name: UpdateCustomer :one
UPDATE customer SET
name = $1,
updated_at = $2
WHERE id = $3
RETURNING *;
