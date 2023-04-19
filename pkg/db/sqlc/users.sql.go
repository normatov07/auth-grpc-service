// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: users.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  role_id, email,password_hash,active
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, account_id, role_id, email, password_hash, password_changed_at, active, login_at, created_at
`

type CreateUserParams struct {
	RoleID       int64  `json:"role_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Active       bool   `json:"active"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.RoleID,
		arg.Email,
		arg.PasswordHash,
		arg.Active,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.RoleID,
		&i.Email,
		&i.PasswordHash,
		&i.PasswordChangedAt,
		&i.Active,
		&i.LoginAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, account_id, role_id, email, password_hash, password_changed_at, active, login_at, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.RoleID,
		&i.Email,
		&i.PasswordHash,
		&i.PasswordChangedAt,
		&i.Active,
		&i.LoginAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET login_at=$1
WHERE id = $2
`

type UpdateUserParams struct {
	LoginAt time.Time `json:"login_at"`
	ID      int64     `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.LoginAt, arg.ID)
	return err
}