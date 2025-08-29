-- name: CreateUser :one
insert into users (email, name, user_name, password_hash) values ($1, $2, $3, $4) returning *;

-- name: GetUser :one
select * from users where user_id=$1;

-- name: GetUsers :many
select * from users;

-- name: CreateTeam :one
insert into teams(team_id, title, leader_id, descr) values ($1, $2, $3, $4) returning *;

-- name: AddTeamMember :exec
insert into team_members(team_id, user_id) values ($1, $2);

-- name: AddTeamLeader :exec
update teams set leader_id=$2 where teams.team_id=$1;
