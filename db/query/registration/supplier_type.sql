-- name: GetSupplierType :one
SELECT * FROM registration.supplier_type
WHERE id = $1;

-- name: ListSupplierTypes :many
SELECT * FROM registration.supplier_type
ORDER BY description
LIMIT $1 OFFSET $2;

-- name: CreateSupplierType :one
INSERT INTO registration.supplier_type (
    description
) VALUES (
    $1
)
RETURNING *;

-- name: UpdateSupplierType :one
UPDATE registration.supplier_type
SET
  description = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSupplierType :exec
DELETE FROM registration.supplier_type
WHERE id = $1;
