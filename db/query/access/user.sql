-- name: GetUser :one
SELECT * FROM access.user
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM access.user
ORDER BY login
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO access.user (
    login, email, name, phone, hash, active
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateUser :one
UPDATE access.user
SET
  email = $2
, name = $3
, phone = $4
, hash = $5
, active = $6
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM access.user
WHERE id = $1;
