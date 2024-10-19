-- +goose Up
ALTER TABLE channels ADD COLUMN last_fetched TIMESTAMP DEFAULT NULL;

-- +goose Down
ALTER TABLE channels DROP COLUMN last_fetched;
