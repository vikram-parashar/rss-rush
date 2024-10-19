-- +goose Up
CREATE TABLE channels(
  id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  html_url TEXT,
  xml_url TEXT NOT NULL UNIQUE,
  owner_id UUID NOT NULL REFERENCES users(id) on DELETE CASCADE
);

-- +goose Down
DROP TABLE channels;
