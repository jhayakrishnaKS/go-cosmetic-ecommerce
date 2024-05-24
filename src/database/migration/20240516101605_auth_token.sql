-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refresh_tokens(
    token TEXT PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    expires_at TIMESTAMP WITHOUT TIME ZONE
);
CREATE TABLE IF NOT EXISTS access_token(
    token TEXT PRIMARY KEY,
    refresh_tokens TEXT REFERENCES refresh_tokens(token),
    user_id UUID REFERENCES users(id),
    expires_at TIMESTAMP WITHOUT TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
