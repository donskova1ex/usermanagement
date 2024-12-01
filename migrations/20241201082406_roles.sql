-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles (
  uuid VARCHAR(36),
  id SERIAL PRIMARY KEY,
  name VARCHAR(64),
  activity BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd


