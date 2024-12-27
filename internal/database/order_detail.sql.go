// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: order_detail.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createOrderDetail = `-- name: CreateOrderDetail :one
INSERT INTO order_detail(
    id, created_at, updated_at, 
    order_id, quantity, product_id
    ) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, quantity, order_id, product_id
`

type CreateOrderDetailParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	OrderID   uuid.UUID
	Quantity  int32
	ProductID uuid.UUID
}

func (q *Queries) CreateOrderDetail(ctx context.Context, arg CreateOrderDetailParams) (OrderDetail, error) {
	row := q.db.QueryRowContext(ctx, createOrderDetail,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.OrderID,
		arg.Quantity,
		arg.ProductID,
	)
	var i OrderDetail
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Quantity,
		&i.OrderID,
		&i.ProductID,
	)
	return i, err
}

const createOrderDetailAndUpdateProduct = `-- name: CreateOrderDetailAndUpdateProduct :one
SELECT create_order_detail_and_update_quantity($1, $2, $3)
`

type CreateOrderDetailAndUpdateProductParams struct {
	POrderID   uuid.UUID
	PProductID uuid.UUID
	PQuantity  int32
}

func (q *Queries) CreateOrderDetailAndUpdateProduct(ctx context.Context, arg CreateOrderDetailAndUpdateProductParams) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, createOrderDetailAndUpdateProduct, arg.POrderID, arg.PProductID, arg.PQuantity)
	var create_order_detail_and_update_quantity interface{}
	err := row.Scan(&create_order_detail_and_update_quantity)
	return create_order_detail_and_update_quantity, err
}
