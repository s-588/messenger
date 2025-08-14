package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// AuthRepository defines persistence methods for authentication and authorization.
type AuthRepository interface {
	// Register creates a new user with the given username and password hash.
	// Returns the new user's ID and registration timestamp.
	Register(ctx context.Context, params RegisterParams) (uuid.UUID, time.Time, error)

	// Authorize checks if the token is valid and the user has the given permission.
	Authorize(ctx context.Context, params AuthorizeParams) (bool, error)

	// Authenticate verifies credentials and returns a new JWT token string.
	Authenticate(ctx context.Context, params AuthenticateParams) (string, error)

	// ChangeUsername updates the user's username.
	ChangeUsername(ctx context.Context, params ChangeUsernameParams) error

	// ChangePassword updates the user's password hash.
	ChangePassword(ctx context.Context, params ChangePasswordParams) error

	// DeleteUser removes the user account by ID.
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type RegisterParams struct {
	Username string
	Password string
}

type AuthorizeParams struct {
	Token      string
	Permission int
}

type AuthenticateParams struct {
	Username string
	Password string
}

type ChangeUsernameParams struct {
	UserID   uuid.UUID
	Username string
}

type ChangePasswordParams struct {
	UserID   uuid.UUID
	Password string
}
