-- name: GetUserByID :one
SELECT username ,first_name , last_name , email , avatar 
FROM users 
WHERE user_id = $1;

-- name: GetUserByUsername :one
SELECT user_id, first_name , last_name , email , avatar 
FROM users 
WHERE username = $1;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users(
    username, first_name, last_name, email, avatar,registered_at
) VALUES(
    $1,$2,$3,$4,$5,$6
)
RETURNING *;

-- name: UpdateUserName :exec
UPDATE users
    set username = $2
WHERE user_id = $1;

-- name: UpdateUserFullName :exec
UPDATE users
    set first_name = $2,
        last_name = $3
WHERE user_id = $1;

-- name: UpdateUserEmail :exec
UPDATE users
    set email = $2
WHERE user_id = $1;

-- name: UpdateUserAvatar :exec
UPDATE users
    set avatar = $2
WHERE user_id = $1;

-- name: SearchUsersByUsername :many
SELECT username, first_name, last_name, avatar
FROM users 
WHERE to_tsvector(username) @@ to_tsquery($1);
