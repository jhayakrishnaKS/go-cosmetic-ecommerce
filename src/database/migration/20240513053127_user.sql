-- +goose Up
-- +goose StatementBegin
-- 20240513053127_user.sql
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY,
    email TEXT,
    username TEXT,
    password TEXT,
    role_id UUID REFERENCES roles(id),
    created_at TIMESTAMP WITHOUT TIME ZONE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
