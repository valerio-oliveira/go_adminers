// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: supplier_type.sql

package db

import (
	"context"
)

const createSupplierType = `-- name: CreateSupplierType :one
INSERT INTO registration.supplier_type (
    description
) VALUES (
    $1
)
RETURNING id, description
`

func (q *Queries) CreateSupplierType(ctx context.Context, description string) (RegistrationSupplierType, error) {
	row := q.db.QueryRowContext(ctx, createSupplierType, description)
	var i RegistrationSupplierType
	err := row.Scan(&i.ID, &i.Description)
	return i, err
}

const deleteSupplierType = `-- name: DeleteSupplierType :exec
DELETE FROM registration.supplier_type
WHERE id = $1
`

func (q *Queries) DeleteSupplierType(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteSupplierType, id)
	return err
}

const getSupplierType = `-- name: GetSupplierType :one
SELECT id, description FROM registration.supplier_type
WHERE id = $1
`

func (q *Queries) GetSupplierType(ctx context.Context, id int32) (RegistrationSupplierType, error) {
	row := q.db.QueryRowContext(ctx, getSupplierType, id)
	var i RegistrationSupplierType
	err := row.Scan(&i.ID, &i.Description)
	return i, err
}

const listSupplierTypes = `-- name: ListSupplierTypes :many
SELECT id, description FROM registration.supplier_type
ORDER BY description
LIMIT $1 OFFSET $2
`

type ListSupplierTypesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListSupplierTypes(ctx context.Context, arg ListSupplierTypesParams) ([]RegistrationSupplierType, error) {
	rows, err := q.db.QueryContext(ctx, listSupplierTypes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RegistrationSupplierType
	for rows.Next() {
		var i RegistrationSupplierType
		if err := rows.Scan(&i.ID, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSupplierType = `-- name: UpdateSupplierType :one
UPDATE registration.supplier_type
SET
  description = $2
WHERE id = $1
RETURNING id, description
`

type UpdateSupplierTypeParams struct {
	ID          int32  `json:"id"`
	Description string `json:"description"`
}

func (q *Queries) UpdateSupplierType(ctx context.Context, arg UpdateSupplierTypeParams) (RegistrationSupplierType, error) {
	row := q.db.QueryRowContext(ctx, updateSupplierType, arg.ID, arg.Description)
	var i RegistrationSupplierType
	err := row.Scan(&i.ID, &i.Description)
	return i, err
}
