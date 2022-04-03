-- name: GetCondoType :one
SELECT * FROM registration.condo_type
WHERE id = $1;

-- name: ListCondoTypes :many
SELECT * FROM registration.condo_type
ORDER BY description
LIMIT $1 OFFSET $2;

-- name: CreateCondoType :one
INSERT INTO registration.condo_type (
    description
) VALUES (
    $1
)
RETURNING *;

-- name: UpdateCondoType :one
UPDATE registration.condo_type
SET
  description = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCondoType :exec
DELETE FROM registration.condo_type
WHERE id = $1;
