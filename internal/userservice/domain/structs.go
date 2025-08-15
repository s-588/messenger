package domain

import "github.com/google/uuid"

type UserData struct {
	user_id    uuid.UUID
	first_name string
	last_name  string
	email      string
	avatar_url string
}
