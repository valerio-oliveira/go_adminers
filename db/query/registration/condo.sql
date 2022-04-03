-- name: GetCondo :one
SELECT * FROM registration.condo
WHERE id = $1;

-- name: ListCondos :many
SELECT * FROM registration.condo
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: CreateCondo :one
INSERT INTO registration.condo (
    id_condo_type, name, nickname, address1, address2, phone1, phone2, email, cnpj
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateCondo :one
UPDATE registration.condo
SET
  id_condo_type = $2
, name = $3
, nickname = $4
, address1 = $5
, address2 = $6
, phone1 = $7
, phone2 = $8
, email = $9
, cnpj = $10
WHERE id = $1
RETURNING *;

-- name: DeleteCondo :exec
DELETE FROM registration.condo
WHERE id = $1;
