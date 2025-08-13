-- name: GetMessageByID :one
SELECT conversation_id, sender_id, message_body 
FROM messages
WHERE message_id = $1;

-- name: GetAllMessagesBySenderID :many
SELECT message_id,conversation_id, message_body 
FROM messages
WHERE sender_id = $1
ORDER BY sent_at DESC
LIMIT $2 OFFSET $3;

-- name: GetAllMessagesByConversationID :many
SELECT message_id,conversation_id, message_body 
FROM messages
WHERE conversation_id = $1
ORDER BY sent_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateMessage :one
INSERT INTO messages(
    conversation_id,sender_id, message_body
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: DeleteMessageByID :exec
DELETE FROM messages
WHERE message_id = $1;

-- name: DeleteMessagesBySenderID :exec
DELETE FROM messages
WHERE sender_id = $1;

-- name: UpdateMessageBody :exec
Update messages
    set message_body = $2
WHERE message_id = $1;

-- name: FindMessagesByMessageBody :many
SELECT message_id, conversation_id, sender_id, message_body
FROM messages
WHERE to_tsvector(message_body) @@ to_tsquery($1)
ORDER BY sent_at DESC
LIMIT $2 OFFSET $3;

-- name: GetConversationsByID :one
SELECT name, creation_date 
FROM conversations
WHERE conversation_id = $1;

-- name: CreateConversation :one
INSERT INTO conversations(
    name
)VALUES(
    $1
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

-- name: GetParticipantByID :one
SELECT conversation_id, user_id 
FROM participants
WHERE participant_id = $1;

-- name: GetAllParticipantsByConversationID :many
SELECT participant_id, user_id
FROM participants
WHERE conversation_id = $1;

-- name: GetAllParticipantsByUserID :many
SELECT participant_id, conversation_id 
FROM participants
WHERE user_id = $1;

-- name: CreateParticipant :one
INSERT INTO participants(
    conversation_id, user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: DeleteParticipantByUserID :exec
DELETE FROM participants
WHERE user_id = $1;

-- name: DeleteParticipantByConversationID :exec
DELETE FROM participants
WHERE conversation_id = $1;

-- name: DeleteParticipantByID :exec
DELETE FROM participants
WHERE participant_id = $1;
-- name: GetAttachmentByID :one
SELECT message_id, file_name, file_type, file_size, file_url
FROM attachments 
WHERE attachment_id = $1;

-- name: GetAttachmentByMessageID :one
SELECT attachment_id, file_name, file_type, file_size, file_url
FROM attachments 
WHERE message_id = $1;

-- name: CreateAttachment :one
INSERT INTO attachments(
    message_id, file_name, file_type, file_size, file_url
) VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: DeleteAttachmentByID :exec
DELETE FROM attachments
WHERE attachment_id = $1;

-- name: DeleteAttachmentByMessageID :exec
DELETE FROM attachments
WHERE message_id = $1;


