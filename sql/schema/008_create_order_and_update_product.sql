-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION create_order_detail_and_update_quantity(
    p_order_id UUID, 
    p_product_id UUID,
    p_quantity INT
    )
returns void AS $$
BEGIN
    INSERT INTO order_detail(order_id, product_id, quantity, created_at, updated_at)
    VALUES (p_order_id, p_product_id, p_quantity, NOW(), NOW());

    UPDATE product
    SET quantity = quantity - p_quantity
    WHERE id = p_product_id AND quantity >= p_quantity AND quantity >= p_quantity;

    IF NOT FOUND THEN 
        RAISE EXCEPTION 'insufficient quantity stock for product';
    END IF;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS create_order_detail_and_update_quantity;
-- +goose StatementEnd
