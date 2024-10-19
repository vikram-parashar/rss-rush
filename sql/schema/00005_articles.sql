-- +goose Up
CREATE TABLE articles(
  id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  link TEXT NOT NULL UNIQUE,
  description TEXT,
  pub_date TIMESTAMP NOT NULL,
  channel_id UUID NOT NULL REFERENCES channels(id) on DELETE CASCADE
);

-- +goose Down
DROP TABLE articles;
