-- name: GetBuilding :one
SELECT * FROM registration.building
WHERE id = $1;

-- name: ListBuildings :many
SELECT * FROM registration.building
ORDER BY description
LIMIT $1 OFFSET $2;

-- name: CreateBuilding :one
INSERT INTO registration.building (
    id_condo, id_building_type, description
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateBuilding :one
UPDATE registration.building
SET --SET id_condo,
  id_building_type = $2
, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteBuilding :exec
DELETE FROM registration.building
WHERE id = $1;
