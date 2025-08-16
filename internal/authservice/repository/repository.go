package repository

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/s-588/messenger/internal/authservice/domain"
)

var (
	ErrNotFound                   = errors.New("resource not found")
	ErrDuplicateUsername = errors.New("username already exist")
	ErrInvalidInput           = errors.New("invalid input")
)

// AuthRepository defines methods for authentication, authorization, and user management.
type AuthRepository interface {
	// CreateUser creates a new user with the given username and hashed password.
	CreateUser(ctx context.Context, username, passHash string) (*domain.User, error)

	// DeleteExpiredTokens removes all refresh tokens that have expired.
	DeleteExpiredTokens(ctx context.Context) error

	// DeleteUserByID deletes a user by their UUID.
	DeleteUserByID(ctx context.Context, id uuid.UUID) error

	// GetRefreshTokenByHash retrieves a refresh token by its hashed value.
	GetRefreshTokenByHash(ctx context.Context, tokenHash string) (*domain.RefreshToken, error)

	// GetUserByID retrieves a user by their UUID.
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)

	// GetUserByUsername retrieves a user by their username.
	// Used during authentication to verify credentials.
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)

	// InsertRefreshToken creates a new refresh token for a user with the provided parameters.
	InsertRefreshToken(ctx context.Context, arg InsertRefreshTokenParams) (*domain.RefreshToken, error)

	// ListUserTokens retrieves all refresh tokens for a user, ordered by issued_at (descending).
	ListUserTokens(ctx context.Context, id uuid.UUID) ([]domain.RefreshToken, error)

	// RevokeRefreshToken deletes a refresh token by its UUID.
	RevokeRefreshToken(ctx context.Context, id uuid.UUID) error

	// UpdateUserName updates a user’s username by their UUID.
	UpdateUserName(ctx context.Context, id uuid.UUID, username string) error

	// UpdateUserPassword updates a user’s password hash by their UUID.
	UpdateUserPassword(ctx context.Context, id uuid.UUID, passHash string) error
}

type InsertRefreshTokenParams struct {
	UserID     uuid.UUID // The ID of the user associated with the token.
	TokenHash  string    // The hashed refresh token for secure storage.
	ExpiresAt  time.Time // The expiration timestamp of the token.
	DeviceInfo string    // Device information (e.g., "iPhone, iOS 16, Safari") from User-Agent or client.
	IpAddress  net.IP    // The client’s IP address.
}
