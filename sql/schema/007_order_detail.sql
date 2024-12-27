-- +goose Up
CREATE TABLE order_detail (
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    order_id UUID NOT NULL REFERENCES "order"(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE
);  

-- +goose Down
DROP TABLE order_detail;