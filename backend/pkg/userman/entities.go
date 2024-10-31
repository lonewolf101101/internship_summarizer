package userman

import (
	"database/sql"
	"time"

	"undrakh.net/summarizer/pkg/entities"
	"undrakh.net/summarizer/pkg/roleman"
)

const (
	AUTH_TYPE_BASIC  = "username_password"
	AUTH_TYPE_GOOGLE = "google"
)

type User struct {
	entities.Model
	UUID           string       `json:"uuid"`
	AuthType       string       `json:"auth_type"`
	PasswordHash   string       `json:"-"`
	Name           string       `json:"name"`
	PhoneNumber    string       `json:"phone_number"`
	Email          string       `json:"email"`
	GoogleID       string       `json:"google_id"`
	ProfilePicture string       `json:"profile_picture"`
	IsVerified     bool         `json:"is_verified"`
	LastLogin      time.Time    `json:"last_login"`
	SelfDeletedAt  sql.NullTime `json:"self_deleted_at,omitempty"`
}

type UserRole struct {
	entities.Model
	URID uint   `json:"id"`
	RID  int    `json:"rid"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type UserWithRoles struct {
	User  User           `json:"user"`
	Roles []roleman.Role `json:"roles"`
}
