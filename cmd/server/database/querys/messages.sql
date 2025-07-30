-- name: GetMessageByID :one
SELECT conversation_id, sender_id, message_body 
FROM messages
WHERE message_id = $1;

-- name: GetAllMessagesBySenderID :many
SELECT message_id,conversation_id, message_body 
FROM messages
WHERE sender_id = $1
ORDER BY sended_at DESC
LIMIT $2 OFFSET $3;

-- name: GetAllMessagesByConversationID :many
SELECT message_id,conversation_id, message_body 
FROM messages
WHERE conversation_id = $1
ORDER BY sended_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateMessage :one
INSERT INTO messages(
    conversation_id,sender_id, message_body, sended_at
) VALUES (
    $1, $2, $3, now()
) RETURNING *;

-- name: DeleteMessageByID :exec
DELETE FROM messages
WHERE message_id = $1;

-- name: DeleteMessagesBySenderID :exec
DELETE FROM messages
WHERE sender_id = $1;

-- name: FindMessagesByMessageBody :many
SELECT message_id, conversation_id, sender_id, message_body
FROM messages
WHERE to_tsvector(message_body) @@ to_tsquery($1)
ORDER BY sended_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateMessageBody :exec
Update messages
    set message_body = $2
WHERE message_id = $1;
