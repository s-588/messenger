-- name: GetUserDataByID :one
SELECT first_name , last_name , email , avatar_url
FROM usersData 
WHERE user_id = $1;

-- name: GetUserDataByEmail :one
SELECT user_id, first_name , last_name , avatar_url
FROM usersData 
WHERE email = $1;

-- name: DeleteUserDataByID :exec
DELETE FROM usersData
WHERE user_id = $1;

-- name: CreateUserData :one
INSERT INTO usersData(
    user_id, first_name, last_name, email, avatar_url
) VALUES(
    $1,$2,$3,$4,$5
)
RETURNING *;

-- name: UpdateUserFullName :exec
UPDATE usersData
    set first_name = $2,
        last_name = $3
WHERE user_id = $1;

-- name: UpdateUserEmail :exec
UPDATE usersData
    set email = $2
WHERE user_id = $1;

-- name: UpdateUserAvatar :exec
UPDATE usersData
    set avatar_url = $2
WHERE user_id = $1;
