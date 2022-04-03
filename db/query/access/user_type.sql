-- name: GetUserType :one
SELECT * FROM access.user_type
WHERE id = $1;

-- name: ListUserTypes :many
SELECT * FROM access.user_type
ORDER BY description
LIMIT $1 OFFSET $2;

-- name: CreateUserType :one
INSERT INTO access.user_type (
    description
) VALUES (
    $1
)
RETURNING *;

-- name: UpdateUserType :one
UPDATE access.user_type
SET
  description = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUserType :exec
DELETE FROM access.user_type
WHERE id = $1;
