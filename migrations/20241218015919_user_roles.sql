-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_roles (
    user_uuid VARCHAR REFERENCES users(uuid),
    role_uuid VARCHAR REFERENCES roles(uuid),
    active BOOLEAN,
    PRIMARY KEY (user_uuid, role_uuid));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
-- +goose StatementEnd
