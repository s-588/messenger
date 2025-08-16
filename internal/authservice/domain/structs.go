package domain

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID
	Username     string
	PasswordHash string
	RegisteredAt time.Time
}

type RefreshToken struct {
	TokenID    uuid.UUID
	UserID     uuid.UUID
	TokenHash  string
	IssuedAt   time.Time
	ExpiresAt  time.Time
	DeviceInfo string
	IPAddress  net.IP
}
