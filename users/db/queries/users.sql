-- name: CreateUser :one
insert into users (email, name, user_name, password_hash) values ($1, $2, $3, $4) returning*;

-- name: GetUser :one
select * from users where user_id=$1;
