-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products(
    id UUID PRIMARY KEY,
    product_title TEXT,
    description TEXT,
    price DOUBLE PRECISION,
    brand TEXT
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
