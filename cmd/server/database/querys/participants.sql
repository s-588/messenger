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
