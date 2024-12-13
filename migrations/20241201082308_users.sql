-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  uuid VARCHAR(36) unique,
  id SERIAL PRIMARY KEY,
  userName VARCHAR(64),
  passwordhash VARCHAR,
  firstName VARCHAR(64),
  lastName VARCHAR(64),
  fullName VARCHAR,
  email VARCHAR(64),
  userGroupId VARCHAR,
  active BOOLEAN,
  deleted BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd


