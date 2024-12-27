-- +goose Up
CREATE TABLE "order" (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    customer_id UUID NOT NULL REFERENCES customer(id) ON DELETE CASCADE
);   

-- +goose Down
DROP TABLE "order";