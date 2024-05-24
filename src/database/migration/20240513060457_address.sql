-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS address (
    id UUID PRIMARY KEY,
    door_no numeric,
    street text,
    city text,
    zipcode numeric,
    user_id UUID REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
