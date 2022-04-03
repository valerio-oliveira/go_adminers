-- name: GetMovementType :one
SELECT * FROM finance.movement_type
WHERE id = $1;

-- name: ListMovementTypes :many
SELECT * FROM finance.movement_type
ORDER BY description
LIMIT $1 OFFSET $2;

-- name: CreateMovementType :one
INSERT INTO finance.movement_type (
    description, direction, id_default_supplier
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateMovementType :one
UPDATE finance.movement_type
SET
  description = $2
, direction = $3
, id_default_supplier = $4
WHERE id = $1
RETURNING *;

-- name: DeleteMovementType :exec
DELETE FROM finance.movement_type
WHERE id = $1;
