// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO access.user (
    login, email, name, phone, hash, active
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, login, email, name, phone, hash, active
`

type CreateUserParams struct {
	Login  string         `json:"login"`
	Email  string         `json:"email"`
	Name   string         `json:"name"`
	Phone  sql.NullString `json:"phone"`
	Hash   string         `json:"hash"`
	Active bool           `json:"active"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (AccessUser, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Login,
		arg.Email,
		arg.Name,
		arg.Phone,
		arg.Hash,
		arg.Active,
	)
	var i AccessUser
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Email,
		&i.Name,
		&i.Phone,
		&i.Hash,
		&i.Active,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM access.user
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, login, email, name, phone, hash, active FROM access.user
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (AccessUser, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i AccessUser
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Email,
		&i.Name,
		&i.Phone,
		&i.Hash,
		&i.Active,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, login, email, name, phone, hash, active FROM access.user
ORDER BY login
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]AccessUser, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccessUser
	for rows.Next() {
		var i AccessUser
		if err := rows.Scan(
			&i.ID,
			&i.Login,
			&i.Email,
			&i.Name,
			&i.Phone,
			&i.Hash,
			&i.Active,
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

const updateUser = `-- name: UpdateUser :one
UPDATE access.user
SET
  email = $2
, name = $3
, phone = $4
, hash = $5
, active = $6
WHERE id = $1
RETURNING id, login, email, name, phone, hash, active
`

type UpdateUserParams struct {
	ID     int32          `json:"id"`
	Email  string         `json:"email"`
	Name   string         `json:"name"`
	Phone  sql.NullString `json:"phone"`
	Hash   string         `json:"hash"`
	Active bool           `json:"active"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (AccessUser, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.Name,
		arg.Phone,
		arg.Hash,
		arg.Active,
	)
	var i AccessUser
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Email,
		&i.Name,
		&i.Phone,
		&i.Hash,
		&i.Active,
	)
	return i, err
}
