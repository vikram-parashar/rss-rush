-- name: CreateChannel :one
INSERT INTO channels( id, created_at, updated_at, name,html_url, xml_url,owner_id) VALUES
(gen_random_uuid(), now(), now(), $1, $2, $3, $4)
RETURNING *;

-- name: GetChannels :many
SELECT * FROM channels
ORDER BY last_fetched NULLS FIRST
LIMIT $1 OFFSET $2;

-- name: UpdateFetched :exec
UPDATE channels
SET last_fetched=now()
WHERE id=$1;

-- name: DeleteChannel :exec
DELETE FROM channels
WHERE id=$1 and owner_id=$2;
