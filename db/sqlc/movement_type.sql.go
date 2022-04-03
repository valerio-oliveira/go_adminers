// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: movement_type.sql

package db

import (
	"context"
	"database/sql"
)

const createMovementType = `-- name: CreateMovementType :one
INSERT INTO finance.movement_type (
    description, direction, id_default_supplier
) VALUES (
    $1, $2, $3
)
RETURNING id, description, direction, id_default_supplier
`

type CreateMovementTypeParams struct {
	Description       string        `json:"description"`
	Direction         string        `json:"direction"`
	IDDefaultSupplier sql.NullInt32 `json:"id_default_supplier"`
}

func (q *Queries) CreateMovementType(ctx context.Context, arg CreateMovementTypeParams) (FinanceMovementType, error) {
	row := q.db.QueryRowContext(ctx, createMovementType, arg.Description, arg.Direction, arg.IDDefaultSupplier)
	var i FinanceMovementType
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Direction,
		&i.IDDefaultSupplier,
	)
	return i, err
}

const deleteMovementType = `-- name: DeleteMovementType :exec
DELETE FROM finance.movement_type
WHERE id = $1
`

func (q *Queries) DeleteMovementType(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteMovementType, id)
	return err
}

const getMovementType = `-- name: GetMovementType :one
SELECT id, description, direction, id_default_supplier FROM finance.movement_type
WHERE id = $1
`

func (q *Queries) GetMovementType(ctx context.Context, id int32) (FinanceMovementType, error) {
	row := q.db.QueryRowContext(ctx, getMovementType, id)
	var i FinanceMovementType
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Direction,
		&i.IDDefaultSupplier,
	)
	return i, err
}

const listMovementTypes = `-- name: ListMovementTypes :many
SELECT id, description, direction, id_default_supplier FROM finance.movement_type
ORDER BY description
LIMIT $1 OFFSET $2
`

type ListMovementTypesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMovementTypes(ctx context.Context, arg ListMovementTypesParams) ([]FinanceMovementType, error) {
	rows, err := q.db.QueryContext(ctx, listMovementTypes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FinanceMovementType
	for rows.Next() {
		var i FinanceMovementType
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.Direction,
			&i.IDDefaultSupplier,
		); err != nil {
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

const updateMovementType = `-- name: UpdateMovementType :one
UPDATE finance.movement_type
SET
  description = $2
, direction = $3
, id_default_supplier = $4
WHERE id = $1
RETURNING id, description, direction, id_default_supplier
`

type UpdateMovementTypeParams struct {
	ID                int32         `json:"id"`
	Description       string        `json:"description"`
	Direction         string        `json:"direction"`
	IDDefaultSupplier sql.NullInt32 `json:"id_default_supplier"`
}

func (q *Queries) UpdateMovementType(ctx context.Context, arg UpdateMovementTypeParams) (FinanceMovementType, error) {
	row := q.db.QueryRowContext(ctx, updateMovementType,
		arg.ID,
		arg.Description,
		arg.Direction,
		arg.IDDefaultSupplier,
	)
	var i FinanceMovementType
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Direction,
		&i.IDDefaultSupplier,
	)
	return i, err
}