-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    uuid VARCHAR(36) unique,
    id SERIAL PRIMARY KEY,
    person_id SERIAL REFERENCES persons(id),
    user_group_id SERIAL REFERENCES user_groups(id),
    active BOOLEAN,
    deleted BOOLEAN
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd