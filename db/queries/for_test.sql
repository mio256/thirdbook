-- name: Ping :exec
select 1;

-- name: TestCreateUser :one
insert into users (name, email, password) values ($1, $2, $3) returning *;

-- name: TestDeleteUser :exec
delete from users where id = $1;
