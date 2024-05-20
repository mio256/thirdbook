-- name: CreateBooking :one
insert into bookings (name, start_time, end_time, user_id, status) values ($1, $2, $3, $4, $5) returning *;

-- name: DeleteBooking :exec
delete from bookings where id = $1;

-- name: GetBooking :one
select * from bookings where id = $1;

-- name: GetBookingStatus :one
select status from bookings where id = $1;

-- name: GetBookings :many
select * from bookings where status = $1 and user_id = $2 and (end_time >= @start_limit or start_time <= @end_limit);

-- name: UpdateBookingStatus :one
update bookings set status = $1 where id = $2 returning *;
