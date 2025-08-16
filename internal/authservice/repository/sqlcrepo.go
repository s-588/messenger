package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/s-588/messenger/internal/authservice/domain"
	postgres "github.com/s-588/messenger/internal/authservice/repository/sqlc"
)

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserAlreadyExist     = errors.New("user already exist")
	ErrRequiredFieldMissing = errors.New("requeired field missing")
)

type SQLCRepo struct {
	queries *postgres.Queries
}

func parseError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return ErrUserNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return ErrUserAlreadyExist
		case pgerrcode.NotNullViolation:
			return ErrRequiredFieldMissing
		}
	}
	slog.Error("unknown db error", "error", err)
	return fmt.Errorf("db error: %w", err)
}

func (repo *SQLCRepo) CreateUser(ctx context.Context, username, passHash string) (*domain.User, error) {
	row, err := repo.queries.CreateUser(ctx, postgres.CreateUserParams{
		Username:     username,
		PasswordHash: passHash,
	})
	return &domain.User{
		UserID:       row.UserID,
		Username:     username,
		RegisteredAt: row.RegisteredAt.Time,
	}, parseError(err)
}

// DeleteExpiredTokens removes all refresh tokens that have expired.
func (repo *SQLCRepo) DeleteExpiredTokens(ctx context.Context) error {
	return parseError(repo.queries.DeleteExpiredTokens(ctx))
}

// DeleteUserByID deletes a user by their UUID.
func (repo *SQLCRepo) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	return parseError(repo.queries.DeleteUserByID(ctx, id))
}

// GetRefreshTokenByHash retrieves a refresh token by its hashed value.
func (repo *SQLCRepo) GetRefreshTokenByHash(ctx context.Context, tokenHash string) (*domain.RefreshToken, error) {
	row, err := repo.queries.GetRefreshTokenByHash(ctx, tokenHash)
	return &domain.RefreshToken{
		TokenID:    row.TokenID,
		UserID:     row.UserID,
		TokenHash:  tokenHash,
		IssuedAt:   row.IssuedAt.Time,
		ExpiresAt:  row.ExpiresAt.Time,
		DeviceInfo: row.DeviceInfo.String,
		IPAddress:  net.IP(row.IpAddress.String),
	}, parseError(err)
}

// GetUserByID retrieves a user by their UUID.
func (repo *SQLCRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row, err := repo.queries.GetUserByID(ctx, id)
	return &domain.User{
		UserID:       id,
		Username:     row.Username,
		RegisteredAt: row.RegisteredAt.Time,
	}, parseError(err)

}

// GetUserByUsername retrieves a user by their username.
// Used during authentication to verify credentials.
func (repo *SQLCRepo) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	row, err := repo.queries.GetUserByUsername(ctx, username)
	return &domain.User{
		UserID:       row.UserID,
		Username:     username,
		PasswordHash: row.PasswordHash,
		RegisteredAt: row.RegisteredAt.Time,
	}, parseError(err)

}

// InsertRefreshToken creates a new refresh token for a user with the provided parameters.
func (repo *SQLCRepo) InsertRefreshToken(ctx context.Context, arg InsertRefreshTokenParams) (*domain.RefreshToken, error) {
	row, err := repo.queries.InsertRefreshToken(ctx, postgres.InsertRefreshTokenParams{
		UserID:    arg.UserID,
		TokenHash: arg.TokenHash,
		ExpiresAt: pgtype.Timestamp{
			Time:  arg.ExpiresAt,
			Valid: true,
		},
		DeviceInfo: pgtype.Text{
			String: arg.DeviceInfo,
			Valid:  true,
		},
		IpAddress: pgtype.Text{
			String: arg.IpAddress.String(),
			Valid:  true,
		},
	})
	return &domain.RefreshToken{
		TokenID:    row.TokenID,
		UserID:     row.TokenID,
		TokenHash:  arg.TokenHash,
		IssuedAt:   row.IssuedAt.Time,
		ExpiresAt:  arg.ExpiresAt,
		DeviceInfo: row.DeviceInfo.String,
		IPAddress:  net.IP(row.IpAddress.String),
	}, parseError(err)

}

// ListUserTokens retrieves all refresh tokens for a user, ordered by issued_at (descending).
func (repo *SQLCRepo) ListUserTokens(ctx context.Context, id uuid.UUID) ([]domain.RefreshToken, error) {
	rows, err := repo.queries.ListUserTokens(ctx, id)
	if err != nil {
		return nil, parseError(err)
	}
	result := make([]domain.RefreshToken, len(rows))

	for i := range rows {
		result[i] = domain.RefreshToken{
			TokenID:    rows[i].TokenID,
			UserID:     id,
			TokenHash:  rows[i].TokenHash,
			IssuedAt:   rows[i].IssuedAt.Time,
			ExpiresAt:  rows[i].ExpiresAt.Time,
			DeviceInfo: rows[i].DeviceInfo.String,
			IPAddress:  net.IP(rows[i].IpAddress.String),
		}
	}

	return result, nil
}

// RevokeRefreshToken deletes a refresh token by its UUID.
func (repo *SQLCRepo) RevokeRefreshToken(ctx context.Context, id uuid.UUID) error {
	return parseError(repo.queries.RevokeRefreshToken(ctx, id))
}

// UpdateUserName updates a user’s username by their UUID.
func (repo *SQLCRepo) UpdateUserName(ctx context.Context, id uuid.UUID, username string) error {
	return parseError(repo.queries.UpdateUserName(ctx, postgres.UpdateUserNameParams{
		UserID:   id,
		Username: username,
	}))
}

// UpdateUserPassword updates a user’s password hash by their UUID.
func (repo *SQLCRepo) UpdateUserPassword(ctx context.Context, id uuid.UUID, passHash string) error {
	return parseError(repo.queries.UpdateUserPassword(ctx, postgres.UpdateUserPasswordParams{
		UserID:       id,
		PasswordHash: passHash,
	}))
}
