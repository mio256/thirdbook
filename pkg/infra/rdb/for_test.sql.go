// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: for_test.sql

package rdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const ping = `-- name: Ping :exec
select 1
`

func (q *Queries) Ping(ctx context.Context) error {
	_, err := q.db.Exec(ctx, ping)
	return err
}

const testCreateBooking = `-- name: TestCreateBooking :one
insert into bookings (name, start_time, end_time, user_id, status) values ($1, $2, $3, $4, $5) returning id, name, start_time, end_time, user_id, status, created_at, updated_at
`

type TestCreateBookingParams struct {
	Name      string           `json:"name"`
	StartTime pgtype.Timestamp `json:"start_time"`
	EndTime   pgtype.Timestamp `json:"end_time"`
	UserID    int64            `json:"user_id"`
	Status    BookingType      `json:"status"`
}

func (q *Queries) TestCreateBooking(ctx context.Context, arg TestCreateBookingParams) (Booking, error) {
	row := q.db.QueryRow(ctx, testCreateBooking,
		arg.Name,
		arg.StartTime,
		arg.EndTime,
		arg.UserID,
		arg.Status,
	)
	var i Booking
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.StartTime,
		&i.EndTime,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const testCreateUser = `-- name: TestCreateUser :one
insert into users (name, email, password) values ($1, $2, $3) returning id, name, email, password, created_at, updated_at
`

type TestCreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) TestCreateUser(ctx context.Context, arg TestCreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, testCreateUser, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const testDeleteBooking = `-- name: TestDeleteBooking :exec
delete from bookings where id = $1
`

func (q *Queries) TestDeleteBooking(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, testDeleteBooking, id)
	return err
}

const testDeleteUser = `-- name: TestDeleteUser :exec
delete from users where id = $1
`

func (q *Queries) TestDeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, testDeleteUser, id)
	return err
}
