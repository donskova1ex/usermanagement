-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_permissions (
  uuid VARCHAR(36),
  AccessPermissionsId SERIAL PRIMARY KEY,
  Name VARCHAR(64),
  Read BOOLEAN,
  Change BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access_permissions;
-- +goose StatementEnd
