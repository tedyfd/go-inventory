// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: customer.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customer(
    id, created_at, updated_at, name
    ) 
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, name
`

type CreateCustomerParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM customer WHERE id=$1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :many
SELECT id, created_at, updated_at, name FROM customer
`

func (q *Queries) GetCustomer(ctx context.Context) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, getCustomer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCustomerByID = `-- name: GetCustomerByID :one
SELECT id, created_at, updated_at, name FROM customer WHERE id = $1
`

func (q *Queries) GetCustomerByID(ctx context.Context, id uuid.UUID) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByID, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const getCustomerByName = `-- name: GetCustomerByName :one
SELECT id, created_at, updated_at, name FROM customer WHERE name = $1
`

func (q *Queries) GetCustomerByName(ctx context.Context, name string) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByName, name)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE customer SET
name = $1,
updated_at = $2
WHERE id = $3
RETURNING id, created_at, updated_at, name
`

type UpdateCustomerParams struct {
	Name      string
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomer, arg.Name, arg.UpdatedAt, arg.ID)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}
