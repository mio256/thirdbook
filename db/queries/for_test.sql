-- name: Ping :exec
select 1;

-- name: TestCreateUser :one
insert into users (name, email, password) values ($1, $2, $3) returning *;

-- name: TestDeleteUser :exec
delete from users where id = $1;

-- name: TestCreateBooking :one
insert into bookings (name, start_time, end_time, user_id, status) values ($1, $2, $3, $4, $5) returning *;

-- name: TestDeleteBooking :exec
delete from bookings where id = $1;
