-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orderStatus(
    id UUID PRIMARY KEY,
    "status" TEXT
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
