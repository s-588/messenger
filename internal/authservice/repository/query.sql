-- name: GetUserByID :one
SELECT username 
FROM users 
WHERE user_id = $1;

-- name: GetUserByUsername :one
SELECT user_id 
FROM users 
WHERE username = $1;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users(
    username, password
) VALUES(
    $1,$2
)
RETURNING user_id;

-- name: UpdateUserName :exec
UPDATE users
    set username = $2
WHERE user_id = $1;

-- name: UpdatePassword :exec
UPDATE users
    set password = $2
WHERE user_id = $1;
