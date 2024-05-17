-- name: CreateUser :one
insert into users (name, email, password) values ($1, $2, $3) returning *;

-- name: DeleteUser :exec
delete from users where id = $1;
