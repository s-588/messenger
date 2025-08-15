package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// AuthRepository defines methods for authentication, authorization and user management.
type AuthRepository interface {
	// Register creates a new user with the given username and password hash.
	// Returns the new user's ID and registration timestamp.
	Register(ctx context.Context, username, pass string) (uuid.UUID, time.Time, error)

	// Authorize checks if the token is valid and the user has the given permission.
	Authorize(ctx context.Context, token string, permission int) (bool, error)

	// Authenticate verifies credentials and returns a new JWT token string.
	Authenticate(ctx context.Context, username, password string) (string, error)

	// ChangeUsername updates the user's username.
	ChangeUsername(ctx context.Context, id uuid.UUID, username string) error

	// ChangePassword updates the user's password hash.
	ChangePassword(ctx context.Context, id uuid.UUID, password string) error

	// DeleteUser removes the user account by ID.
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
