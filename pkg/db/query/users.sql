-- name: CreateUser :one
INSERT INTO users (
  role_id, email,password_hash,active
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users SET login_at=$1
WHERE id = $2;

