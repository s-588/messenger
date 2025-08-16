package domain

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	conversation_id uuid.UUID
	name            string
	creation_date   time.Time
}

type Participant struct {
	participant_id  uuid.UUID
	conversation_id uuid.UUID
	user_id         uuid.UUID
}

type Message struct {
	message_id      uuid.UUID
	conversation_id uuid.UUID
	sender_id       uuid.UUID
	sent_at         time.Time
	message_body    string
}

type Attachment struct {
	attachment_id uuid.UUID
	message_id    uuid.UUID
	file_name     string
	file_type     string
	file_size     int
	file_url      string
}
