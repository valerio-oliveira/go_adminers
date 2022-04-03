// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user__condo.sql

package db

import (
	"context"
)

const createUserCondo = `-- name: CreateUserCondo :one
INSERT INTO access.user__condo (
    id_user, id_condo, id_user_type
) VALUES (
    $1, $2, $3
)
RETURNING id_user, id_condo, id_user_type
`

type CreateUserCondoParams struct {
	IDUser     int32 `json:"id_user"`
	IDCondo    int32 `json:"id_condo"`
	IDUserType int32 `json:"id_user_type"`
}

func (q *Queries) CreateUserCondo(ctx context.Context, arg CreateUserCondoParams) (AccessUserCondo, error) {
	row := q.db.QueryRowContext(ctx, createUserCondo, arg.IDUser, arg.IDCondo, arg.IDUserType)
	var i AccessUserCondo
	err := row.Scan(&i.IDUser, &i.IDCondo, &i.IDUserType)
	return i, err
}

const deleteUserCondo = `-- name: DeleteUserCondo :exec
DELETE FROM access.user__condo
WHERE id_user = $1 AND id_condo = $2
`

type DeleteUserCondoParams struct {
	IDUser  int32 `json:"id_user"`
	IDCondo int32 `json:"id_condo"`
}

func (q *Queries) DeleteUserCondo(ctx context.Context, arg DeleteUserCondoParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserCondo, arg.IDUser, arg.IDCondo)
	return err
}

const getUserCondo = `-- name: GetUserCondo :one
SELECT id_user, id_condo, id_user_type FROM access.user__condo
WHERE id_user = $1 AND id_condo = $2
`

type GetUserCondoParams struct {
	IDUser  int32 `json:"id_user"`
	IDCondo int32 `json:"id_condo"`
}

func (q *Queries) GetUserCondo(ctx context.Context, arg GetUserCondoParams) (AccessUserCondo, error) {
	row := q.db.QueryRowContext(ctx, getUserCondo, arg.IDUser, arg.IDCondo)
	var i AccessUserCondo
	err := row.Scan(&i.IDUser, &i.IDCondo, &i.IDUserType)
	return i, err
}

const listUserCondos = `-- name: ListUserCondos :many
SELECT id_user, id_condo, id_user_type FROM access.user__condo
LIMIT $1 OFFSET $2
`

type ListUserCondosParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUserCondos(ctx context.Context, arg ListUserCondosParams) ([]AccessUserCondo, error) {
	rows, err := q.db.QueryContext(ctx, listUserCondos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccessUserCondo
	for rows.Next() {
		var i AccessUserCondo
		if err := rows.Scan(&i.IDUser, &i.IDCondo, &i.IDUserType); err != nil {
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

const updateUserCondo = `-- name: UpdateUserCondo :one
UPDATE access.user__condo
SET
  id_user_type = $3
WHERE id_user = $1 AND id_condo = $2
RETURNING id_user, id_condo, id_user_type
`

type UpdateUserCondoParams struct {
	IDUser     int32 `json:"id_user"`
	IDCondo    int32 `json:"id_condo"`
	IDUserType int32 `json:"id_user_type"`
}

func (q *Queries) UpdateUserCondo(ctx context.Context, arg UpdateUserCondoParams) (AccessUserCondo, error) {
	row := q.db.QueryRowContext(ctx, updateUserCondo, arg.IDUser, arg.IDCondo, arg.IDUserType)
	var i AccessUserCondo
	err := row.Scan(&i.IDUser, &i.IDCondo, &i.IDUserType)
	return i, err
}
