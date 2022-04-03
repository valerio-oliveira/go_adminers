// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: condo.sql

package db

import (
	"context"
	"database/sql"
)

const createCondo = `-- name: CreateCondo :one
INSERT INTO registration.condo (
    id_condo_type, name, nickname, address1, address2, phone1, phone2, email, cnpj
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, id_condo_type, name, nickname, address1, address2, phone1, phone2, email, cnpj
`

type CreateCondoParams struct {
	IDCondoType int32          `json:"id_condo_type"`
	Name        string         `json:"name"`
	Nickname    string         `json:"nickname"`
	Address1    sql.NullString `json:"address1"`
	Address2    sql.NullString `json:"address2"`
	Phone1      sql.NullString `json:"phone1"`
	Phone2      sql.NullString `json:"phone2"`
	Email       string         `json:"email"`
	Cnpj        int64          `json:"cnpj"`
}

func (q *Queries) CreateCondo(ctx context.Context, arg CreateCondoParams) (RegistrationCondo, error) {
	row := q.db.QueryRowContext(ctx, createCondo,
		arg.IDCondoType,
		arg.Name,
		arg.Nickname,
		arg.Address1,
		arg.Address2,
		arg.Phone1,
		arg.Phone2,
		arg.Email,
		arg.Cnpj,
	)
	var i RegistrationCondo
	err := row.Scan(
		&i.ID,
		&i.IDCondoType,
		&i.Name,
		&i.Nickname,
		&i.Address1,
		&i.Address2,
		&i.Phone1,
		&i.Phone2,
		&i.Email,
		&i.Cnpj,
	)
	return i, err
}

const deleteCondo = `-- name: DeleteCondo :exec
DELETE FROM registration.condo
WHERE id = $1
`

func (q *Queries) DeleteCondo(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCondo, id)
	return err
}

const getCondo = `-- name: GetCondo :one
SELECT id, id_condo_type, name, nickname, address1, address2, phone1, phone2, email, cnpj FROM registration.condo
WHERE id = $1
`

func (q *Queries) GetCondo(ctx context.Context, id int32) (RegistrationCondo, error) {
	row := q.db.QueryRowContext(ctx, getCondo, id)
	var i RegistrationCondo
	err := row.Scan(
		&i.ID,
		&i.IDCondoType,
		&i.Name,
		&i.Nickname,
		&i.Address1,
		&i.Address2,
		&i.Phone1,
		&i.Phone2,
		&i.Email,
		&i.Cnpj,
	)
	return i, err
}

const listCondos = `-- name: ListCondos :many
SELECT id, id_condo_type, name, nickname, address1, address2, phone1, phone2, email, cnpj FROM registration.condo
ORDER BY name
LIMIT $1 OFFSET $2
`

type ListCondosParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCondos(ctx context.Context, arg ListCondosParams) ([]RegistrationCondo, error) {
	rows, err := q.db.QueryContext(ctx, listCondos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RegistrationCondo
	for rows.Next() {
		var i RegistrationCondo
		if err := rows.Scan(
			&i.ID,
			&i.IDCondoType,
			&i.Name,
			&i.Nickname,
			&i.Address1,
			&i.Address2,
			&i.Phone1,
			&i.Phone2,
			&i.Email,
			&i.Cnpj,
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

const updateCondo = `-- name: UpdateCondo :one
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
RETURNING id, id_condo_type, name, nickname, address1, address2, phone1, phone2, email, cnpj
`

type UpdateCondoParams struct {
	ID          int32          `json:"id"`
	IDCondoType int32          `json:"id_condo_type"`
	Name        string         `json:"name"`
	Nickname    string         `json:"nickname"`
	Address1    sql.NullString `json:"address1"`
	Address2    sql.NullString `json:"address2"`
	Phone1      sql.NullString `json:"phone1"`
	Phone2      sql.NullString `json:"phone2"`
	Email       string         `json:"email"`
	Cnpj        int64          `json:"cnpj"`
}

func (q *Queries) UpdateCondo(ctx context.Context, arg UpdateCondoParams) (RegistrationCondo, error) {
	row := q.db.QueryRowContext(ctx, updateCondo,
		arg.ID,
		arg.IDCondoType,
		arg.Name,
		arg.Nickname,
		arg.Address1,
		arg.Address2,
		arg.Phone1,
		arg.Phone2,
		arg.Email,
		arg.Cnpj,
	)
	var i RegistrationCondo
	err := row.Scan(
		&i.ID,
		&i.IDCondoType,
		&i.Name,
		&i.Nickname,
		&i.Address1,
		&i.Address2,
		&i.Phone1,
		&i.Phone2,
		&i.Email,
		&i.Cnpj,
	)
	return i, err
}