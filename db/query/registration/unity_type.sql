-- name: GetUnityType :one
SELECT * FROM registration.unity_type
WHERE id = $1;

-- name: ListUnityTypes :many
SELECT * FROM registration.unity_type
ORDER BY description
LIMIT $1 OFFSET $2;

-- name: CreateUnityType :one
INSERT INTO registration.unity_type (
    description
) VALUES (
    $1
)
RETURNING *;

-- name: UpdateUnityType :one
UPDATE registration.unity_type
SET
  description = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUnityType :exec
DELETE FROM registration.unity_type
WHERE id = $1;
