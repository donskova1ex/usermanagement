-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_roles (
  uuid VARCHAR(36),
  UserRolesid SERIAL PRIMARY KEY,
  UserId VARCHAR,
  RoleId VARCHAR,
  Active BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
-- +goose StatementEnd