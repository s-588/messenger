-- name: GetAttachmentByID :one
SELECT message_id, file_name, file_type, file_size, file_path
FROM attachments 
WHERE attachment_id = $1;

-- name: GetAttachmentByMessageID :one
SELECT attachment_id, file_name, file_type, file_size, file_path
FROM attachments 
WHERE message_id = $1;

-- name: CreateAttachment :one
INSERT INTO attachments(
    message_id, file_name, file_type, file_size, file_path
) VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: DeleteAttachmentByID :exec
DELETE FROM attachments
WHERE attachment_id = $1;

-- name: DeleteAttachmentByMessageID :exec
DELETE FROM attachments
WHERE message_id = $1;
