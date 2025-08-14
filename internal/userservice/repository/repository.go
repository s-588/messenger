package repository

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/s-588/messenger/internal/genproto/userservice"
)

type UserRepository interface {
	GetUserData(ctx context.Context, id uuid.UUID) (*pb.GetUserDataResponse, error)
	DeleteUserData(ctx context.Context, id uuid.UUID) error
}
