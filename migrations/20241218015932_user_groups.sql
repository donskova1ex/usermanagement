-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_groups (
    uuid VARCHAR(36) unique,
    id SERIAL PRIMARY KEY,
    name VARCHAR(36),
    role_uuid VARCHAR REFERENCES roles(uuid),
    user_uuid VARCHAR REFERENCES users(uuid),
    isActive BOOLEAN,
    isDeleted BOOLEAN);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_groups;
-- +goose StatementEnd
