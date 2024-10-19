-- +goose Up
CREATE TABLE users(
  id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  api_key VARCHAR(64) NOT NULL DEFAULT encode(sha256(((random())::text)::bytea), 'hex'::text)
);

-- +goose Down
DROP TABLE users;
