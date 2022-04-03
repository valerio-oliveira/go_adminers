-- name: GetBuildingType :one
SELECT * FROM registration.building_type
WHERE id = $1;

-- name: ListBuildingTypes :many
SELECT * FROM registration.building_type
ORDER BY description
LIMIT $1 OFFSET $2;

-- name: CreateBuildingType :one
INSERT INTO registration.building_type (
    description
) VALUES (
    $1
)
RETURNING *;

-- name: UpdateBuildingType :one
UPDATE registration.building_type
SET
  description = $2
WHERE id = $1
RETURNING *;

-- name: DeleteBuildingType :exec
DELETE FROM registration.building_type
WHERE id = $1;
