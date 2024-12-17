-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS persons (
    uuid VARCHAR(36) unique,
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(32),
    last_name VARCHAR(32),
    full_name VARCHAR(64),
    email VARCHAR(64));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS persons;
-- +goose StatementEnd
