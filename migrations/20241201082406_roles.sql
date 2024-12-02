-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles (
  uuid VARCHAR(36),
  Roleid SERIAL PRIMARY KEY,
  AccessPermissionsId VARCHAR,
  activity BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd