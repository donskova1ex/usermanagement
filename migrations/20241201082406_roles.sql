-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles (
    uuid VARCHAR(36) unique,
    id SERIAL PRIMARY KEY,
    name VARCHAR(36),
    reading BOOLEAN,
    changing BOOLEAN,
    viewing BOOLEAN,
    activity BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd