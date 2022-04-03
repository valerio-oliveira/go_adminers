-- name: GetUnity :one
SELECT * FROM registration.unity
WHERE id = $1;

-- name: ListUnitys :many
SELECT * FROM registration.unity
ORDER BY unity_number
LIMIT $1 OFFSET $2;

-- name: CreateUnity :one
INSERT INTO registration.unity (
    id_building, id_unity_type, unity_number
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateUnity :one
UPDATE registration.unity
SET
  id_building = $2
, id_unity_type = $3
, unity_number = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUnity :exec
DELETE FROM registration.unity
WHERE id = $1;
