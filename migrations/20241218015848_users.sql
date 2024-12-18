-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    uuid VARCHAR(36) unique,
    id SERIAL PRIMARY KEY,
    userName VARCHAR(32),
    passwordhash VARCHAR(64),
    person_uuid VARCHAR(36) REFERENCES persons(uuid),
    active BOOLEAN,
    deleted BOOLEAN);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd