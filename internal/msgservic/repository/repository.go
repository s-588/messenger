package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/s-588/messenger/internal/msgservic/domain"
)

// MsgRepository defines methods for managing messages, attachments, conversations, and participants.
type MsgRepository interface {
	// CreateAttachment adds a new attachment to a message.
	CreateAttachment(ctx context.Context, arg CreateAttachmentParams) (*domain.Attachment, error)

	// CreateConversation creates a new conversation with the given name.
	CreateConversation(ctx context.Context, name string) (*domain.Conversation, error)

	// CreateMessage creates a new message in a conversation by a sender.
	CreateMessage(ctx context.Context, convID, senderID uuid.UUID, msg string) (*domain.Message, error)

	// CreateParticipant adds a user to a conversation as a participant.
	CreateParticipant(ctx context.Context, convID, userID uuid.UUID) (*domain.Participant, error)

	// DeleteAttachmentByID deletes an attachment by its UUID.
	DeleteAttachmentByID(ctx context.Context, attachmentID uuid.UUID) error

	// DeleteAttachmentByMessageID deletes all attachments for a given message.
	DeleteAttachmentByMessageID(ctx context.Context, messageID uuid.UUID) error

	// DeleteConversation deletes a conversation by its UUID.
	DeleteConversation(ctx context.Context, conversationID uuid.UUID) error

	// DeleteMessageByID deletes a message by its UUID.
	DeleteMessageByID(ctx context.Context, messageID uuid.UUID) error

	// DeleteMessagesBySenderID deletes all messages sent by a user.
	DeleteMessagesBySenderID(ctx context.Context, senderID uuid.UUID) error

	// DeleteParticipantByConversationID removes all participants from a conversation.
	DeleteParticipantByConversationID(ctx context.Context, conversationID uuid.UUID) error

	// DeleteParticipantByID deletes a specific participant by their UUID.
	DeleteParticipantByID(ctx context.Context, participantID uuid.UUID) error

	// DeleteParticipantByUserID removes a user from all conversations they participate in.
	DeleteParticipantByUserID(ctx context.Context, userID uuid.UUID) error

	// FindMessagesByMessageBody searches messages using a full-text search query.
	// Returns a list of matching messages up to limit, starting from offset.
	FindMessagesByMessageBody(ctx context.Context, query string, limit, offset int) ([]domain.Message, error)

	// GetAllMessagesByConversationID retrieves messages in a conversation in descending order by sent_at.
	// Returns a paginated list of messages up to limit, starting from offset.
	GetAllMessagesByConversationID(ctx context.Context, convID uuid.UUID, limit, offset int) ([]domain.Message, error)

	// GetAllMessagesBySenderID retrieves messages sent by a user in descending order by sent_at.
	// Returns a paginated list of messages up to limit, starting from offset.
	GetAllMessagesBySenderID(ctx context.Context, senderID uuid.UUID, limit, offset int) ([]domain.Message, error)

	// GetAllParticipantsByConversationID retrieves all participants in a conversation.
	GetAllParticipantsByConversationID(ctx context.Context, conversationID uuid.UUID) ([]domain.Participant, error)

	// GetAllParticipantsByUserID retrieves all conversations a user participates in.
	GetAllParticipantsByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Participant, error)

	// GetAttachmentByID retrieves an attachment by its UUID.
	GetAttachmentByID(ctx context.Context, attachmentID uuid.UUID) (*domain.Attachment, error)

	// GetAttachmentByMessageID retrieves an attachment for a specific message.
	GetAttachmentByMessageID(ctx context.Context, messageID uuid.UUID) (*domain.Attachment, error)

	// GetConversationsByID retrieves a conversation by its UUID.
	GetConversationsByID(ctx context.Context, conversationID uuid.UUID) (*domain.Conversation, error)

	// GetMessageByID retrieves a message by its UUID.
	GetMessageByID(ctx context.Context, messageID uuid.UUID) (*domain.Message, error)

	// GetParticipantByID retrieves a participant by their UUID.
	GetParticipantByID(ctx context.Context, participantID uuid.UUID) (*domain.Participant, error)

	// SearchConversationByName searches conversations using a full-text search query.
	SearchConversationByName(ctx context.Context, toTsquery string) ([]domain.Conversation, error)

	// UpdateConversationName updates a conversation’s name by its UUID.
	UpdateConversationName(ctx context.Context, convID uuid.UUID, name string) error

	// UpdateMessageBody updates a message’s body by its UUID.
	UpdateMessageBody(ctx context.Context, msgID uuid.UUID, msgBody string) error
}

type CreateAttachmentParams struct {
	MessageID uuid.UUID // The ID of the message the attachment belongs to.
	FileName  string    // The name of the attached file.
	FileType  string    // The MIME type of the file (e.g., "image/png").
	FileSize  int32     // The size of the file in bytes.
	FileURL   string    // The URL where the file is stored.
}
