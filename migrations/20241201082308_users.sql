-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  uuid VARCHAR(36),
  id SERIAL PRIMARY KEY,
  name VARCHAR(64),
  passwordhash VARCHAR,
  activity BOOLEAN,
  deleted BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd


