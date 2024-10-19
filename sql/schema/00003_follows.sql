-- +goose Up
CREATE TABLE follows (
    user_id UUID NOT NULL REFERENCES users(id)  on DELETE CASCADE,
    channel_id UUID NOT NULL REFERENCES channels(id)  on DELETE CASCADE,
    PRIMARY KEY (user_id, channel_id)
);

-- +goose Down
DROP TABLE follows;
