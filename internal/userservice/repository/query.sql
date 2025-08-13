-- name: GetUserByID :one
SELECT first_name , last_name , email , avatar_url
FROM users 
WHERE user_id = $1;

-- name: GetUserByEmail :one
SELECT user_id, first_name , last_name , avatar_url
FROM users 
WHERE email = $1;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users(
    user_id, first_name, last_name, email, avatar_url
) VALUES(
    $1,$2,$3,$4,$5
)
RETURNING *;

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
    set avatar_url = $2
WHERE user_id = $1;
