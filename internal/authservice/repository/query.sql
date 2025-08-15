-- name: GetUserByID :one
SELECT user_id, username, registered_at
FROM users 
WHERE user_id = $1;

-- name: GetUserByUsername :one
SELECT user_id, username, registered_at
FROM users 
WHERE username = $1;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users(
    username, password_hash
) VALUES(
    $1,$2
)
RETURNING user_id,registered_at;

-- name: UpdateUserName :exec
UPDATE users
    set username = $2
WHERE user_id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
    set password_hash = $2
WHERE user_id = $1;

-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens (user_id, token_hash, expires_at, device_info, ip_address)
VALUES ($1, $2, $3, $4, $5)
RETURNING token_id, issued_at, expires_at;

-- name: GetRefreshTokenByHash :one
SELECT token_id, user_id, token_hash, issued_at, expires_at
FROM refresh_tokens
WHERE token_hash = $1;

-- name: RevokeRefreshToken :exec
DELETE from refresh_tokens
WHERE token_id = $1;

-- name: DeleteExpiredTokens :exec
DELETE FROM refresh_tokens
WHERE expires_at < now();

-- name: ListUserTokens :many
SELECT token_id, issued_at, expires_at, device_info, ip_address
FROM refresh_tokens
WHERE user_id = $1
ORDER BY issued_at DESC;

