-- name: GetConversationsByID :one
SELECT name, creation_date 
FROM conversations
WHERE conversation_id = $1;

-- name: CreateConversation :one
INSERT INTO conversations(
    name, creation_date
)VALUES(
    $1, now()
)RETURNING *;

-- name: DeleteConversation :exec
DELETE FROM conversations
WHERE conversation_id = $1;

-- name: UpdateConversationName :exec
UPDATE conversations
    set name = $2
WHERE conversation_id = $1;

-- name: SearchConversationByName :many
SELECT conversation_id, name, creation_date
FROM conversations
WHERE to_tsvector(name) @@ to_tsquery($1);
