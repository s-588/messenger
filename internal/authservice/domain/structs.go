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
	TokenID     uuid.UUID
	UserID      string
	TokenHash   string
	IssuedAt    time.Time
	ExpiresAt   time.Time
	device_info string
	ip_address  net.IP
}
