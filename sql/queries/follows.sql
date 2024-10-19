-- name: AddFollow :one
INSERT INTO follows(user_id,channel_id) VALUES
($1, $2)
RETURNING *;

-- name: GetFollows :many
SELECT * FROM follows
WHERE user_id=$1;

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE user_id=$1 and channel_id=$2;
