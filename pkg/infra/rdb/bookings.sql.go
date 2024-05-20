// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: bookings.sql

package rdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBooking = `-- name: CreateBooking :one
insert into bookings (name, start_time, end_time, user_id, status) values ($1, $2, $3, $4, $5) returning id, name, start_time, end_time, user_id, status, created_at, updated_at
`

type CreateBookingParams struct {
	Name      string           `json:"name"`
	StartTime pgtype.Timestamp `json:"start_time"`
	EndTime   pgtype.Timestamp `json:"end_time"`
	UserID    int64            `json:"user_id"`
	Status    BookingType      `json:"status"`
}

func (q *Queries) CreateBooking(ctx context.Context, arg CreateBookingParams) (Booking, error) {
	row := q.db.QueryRow(ctx, createBooking,
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

const deleteBooking = `-- name: DeleteBooking :exec
delete from bookings where id = $1
`

func (q *Queries) DeleteBooking(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteBooking, id)
	return err
}

const getBooking = `-- name: GetBooking :one
select id, name, start_time, end_time, user_id, status, created_at, updated_at from bookings where id = $1
`

func (q *Queries) GetBooking(ctx context.Context, id int64) (Booking, error) {
	row := q.db.QueryRow(ctx, getBooking, id)
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

const getBookingStatus = `-- name: GetBookingStatus :one
select status from bookings where id = $1
`

func (q *Queries) GetBookingStatus(ctx context.Context, id int64) (BookingType, error) {
	row := q.db.QueryRow(ctx, getBookingStatus, id)
	var status BookingType
	err := row.Scan(&status)
	return status, err
}

const getBookings = `-- name: GetBookings :many
select id, name, start_time, end_time, user_id, status, created_at, updated_at from bookings where status = $1 and user_id = $2 and (end_time >= $3 or start_time <= $4)
`

type GetBookingsParams struct {
	Status     BookingType      `json:"status"`
	UserID     int64            `json:"user_id"`
	StartLimit pgtype.Timestamp `json:"start_limit"`
	EndLimit   pgtype.Timestamp `json:"end_limit"`
}

func (q *Queries) GetBookings(ctx context.Context, arg GetBookingsParams) ([]Booking, error) {
	rows, err := q.db.Query(ctx, getBookings,
		arg.Status,
		arg.UserID,
		arg.StartLimit,
		arg.EndLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Booking
	for rows.Next() {
		var i Booking
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.StartTime,
			&i.EndTime,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBookingStatus = `-- name: UpdateBookingStatus :one
update bookings set status = $1 where id = $2 returning id, name, start_time, end_time, user_id, status, created_at, updated_at
`

type UpdateBookingStatusParams struct {
	Status BookingType `json:"status"`
	ID     int64       `json:"id"`
}

func (q *Queries) UpdateBookingStatus(ctx context.Context, arg UpdateBookingStatusParams) (Booking, error) {
	row := q.db.QueryRow(ctx, updateBookingStatus, arg.Status, arg.ID)
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
