package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/s-588/messenger/internal/userservice/domain"
)

// UserRepository defines persistence methods for user profile data.
type UserDataRepository interface {
	// AddUserData creates a new user profile with the provided details.
	AddUserData(ctx context.Context, arg CreateUserParams) (*domain.UserData, error)

	// DeleteUserDataByID deletes a user profile by their UUID.
	DeleteUserDataByID(ctx context.Context, userID uuid.UUID) error

	// GetUserDataByEmail retrieves a user data by their email address.
	GetUserDataByEmail(ctx context.Context, email string) (domain.UserData, error)

	// GetUserDAtaByID retrieves a user profile by their UUID.
	GetUserDataByID(ctx context.Context, userID uuid.UUID) (domain.UserData, error)

	// UpdateUserAvatar updates a user’s avatar URL by their UUID.
	UpdateUserAvatar(ctx context.Context, userID uuid.UUID, avatarURL string) error

	// UpdateUserEmail updates a user’s email address by their UUID.
	UpdateUserEmail(ctx context.Context, userID uuid.UUID, email string) error

	// UpdateUserFullName updates a user’s first and last names by their UUID.
	UpdateUserFullName(ctx context.Context, userID uuid.UUID, firstName, lastName string) error
}

type CreateUserParams struct {
	UserID    uuid.UUID // The unique ID of the user.
	FirstName string    // The user’s first name.
	LastName  string    // The user’s last name, optional.
	Email     string    // The user’s email address, optional.
	AvatarURL string    // The URL of the user’s avatar image, optional.
}
