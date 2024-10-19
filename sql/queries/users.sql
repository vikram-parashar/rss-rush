-- name: CreateUser :one
INSERT INTO users( id, created_at, updated_at, name,email) VALUES
(gen_random_uuid(), now(), now(), $1, $2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE api_key=$1
LIMIT 1;
