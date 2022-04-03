-- name: GetUserCondo :one
SELECT * FROM access.user__condo
WHERE id_user = $1 AND id_condo = $2;

-- name: ListUserCondos :many
SELECT * FROM access.user__condo
LIMIT $1 OFFSET $2;

-- name: CreateUserCondo :one
INSERT INTO access.user__condo (
    id_user, id_condo, id_user_type
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateUserCondo :one
UPDATE access.user__condo
SET
  id_user_type = $3
WHERE id_user = $1 AND id_condo = $2
RETURNING *;

-- name: DeleteUserCondo :exec
DELETE FROM access.user__condo
WHERE id_user = $1 AND id_condo = $2;
