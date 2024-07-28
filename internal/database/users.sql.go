// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
)

const insertUser = `-- name: InsertUser :one
INSERT INTO Users(
    Id,
    FirstName,
    LastName,
    Email,
    Username,
    Password
)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5
)

RETURNING id, firstname, lastname, email, username, password, createdat, updatedat
`

type InsertUserParams struct {
	Firstname string
	Lastname  sql.NullString
	Email     string
	Username  string
	Password  string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Username,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}
