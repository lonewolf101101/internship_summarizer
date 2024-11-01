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
	URID uint   `gorm:"primaryKey;autoIncrement;column:urid" json:"id"`
	RID  int    `gorm:"column:rid" json:"rid"`
	UUID string `gorm:"column:uuid" json:"uuid"`
	Name string `gorm:"column:name" json:"name"`
}

type UserWithRoles struct {
	User  User           `json:"user"`
	Roles []roleman.Role `json:"roles"`
}
