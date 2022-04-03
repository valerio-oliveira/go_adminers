-- name: GetSupplier :one
SELECT * FROM registration.supplier
WHERE id = $1;

-- name: ListSuppliers :many
SELECT * FROM registration.supplier
ORDER BY description
LIMIT $1 OFFSET $2;

-- name: CreateSupplier :one
INSERT INTO registration.supplier (
    id_supplier_type, description
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateSupplier :one
UPDATE registration.supplier
SET
  id_supplier_type = $2
, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteSupplier :exec
DELETE FROM registration.supplier
WHERE id = $1;
