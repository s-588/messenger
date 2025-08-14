package repository

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/s-588/messenger/internal/genproto/msgservice"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MsgRepository interface {
	SendMessage(ctx context.Context, msg *pb.SendMessageRequest) (uuid.UUID, *timestamppb.Timestamp, error)
	GetMessage(ctx context.Context, id uuid.UUID) (*pb.GetMessageResponse, error)
	ListMessages(ctx context.Context, list *pb.ListMessagesRequest) (*pb.ListConversationsResponse, error)
	DeleteMessage(ctx context.Context, id uuid.UUID) error
	EditMessage(ctx context.Context, id uuid.UUID, msg string) error

	GetAttachment(ctx context.Context, id uuid.UUID) (*pb.GetAttachmentResponse, error)
	DeleteAttachments(ctx context.Context, id uuid.UUID) error

	CreateConversation(ctx context.Context, conv *pb.CreateConversationRequest) (uuid.UUID, *timestamppb.Timestamp, error)
	GetConversation(ctx context.Context, id uuid.UUID) (*pb.GetConversationResponse, error)
	ListConversations(ctx context.Context, list *pb.ListConversationsRequest) (*pb.ListConversationsResponse, error)
	AddParticipant(ctx context.Context, userID, conversationID uuid.UUID) error
	DeleteParticipant(ctx context.Context, id uuid.UUID) error
}
