-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_permissions (
  uuid VARCHAR(36),
  id SERIAL PRIMARY KEY,
  name VARCHAR(64),
  viewing BOOLEAN,
  reading BOOLEAN,
  change BOOLEAN );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access_permissions;
-- +goose StatementEnd