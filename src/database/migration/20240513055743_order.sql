-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "order" (
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES products(id),
    user_id UUID REFERENCES users(id),
    orderStatus_id UUID REFERENCES orderStatus(id),
    created_at TIMESTAMP WITHOUT TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
