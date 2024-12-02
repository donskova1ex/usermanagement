-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  uuid VARCHAR(36),
  Userid SERIAL PRIMARY KEY,
  UserName VARCHAR(64),
  passwordhash VARCHAR,
  FirstName VARCHAR(64),
  LastName VARCHAR(64),
  FullName VARCHAR,
  Email VARCHAR(64),
  UserGroupId VARCHAR,
  Active BOOLEAN,
  Deleted BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd


