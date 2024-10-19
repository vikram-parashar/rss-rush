-- name: CreateArticle :one
INSERT INTO articles( id, created_at, updated_at, title, link, description, pub_date, channel_id) VALUES
(gen_random_uuid(), now(), now(), $1, $2, $3, $4, $5)
RETURNING *;

-- name: GetArticles :many
SELECT * FROM articles
INNER JOIN follows on
articles.channel_id = follows.channel_id
WHERE user_id=$1
LIMIT $2 OFFSET $3;

