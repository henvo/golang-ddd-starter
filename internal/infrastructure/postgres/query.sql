-- name: GetUserByUUID :one
SELECT * FROM users
WHERE uuid = $1 LIMIT 1;
