-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  uuid VARCHAR(36) unique,
  id SERIAL PRIMARY KEY,
  login VARCHAR(64),
  passwordhash VARCHAR,
  firstname VARCHAR(64),
  lastname VARCHAR(64),
  fullname VARCHAR,
  email VARCHAR(64),
  active BOOLEAN,
  deleted BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd


