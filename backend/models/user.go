package models

import (
	"github.com/google/uuid"
)

// User contains user information.
type User struct {
	// UUID is the unique identifier of the user.
	UUID uuid.UUID `gorm:"primaryKey"`

	// ID is the ID of the user. One example is the student ID.
	ID string

	// Email is the email address of the user.
	Email string

	// Password is store as plain text.
	//
	// TODO: store password hash and salt.
	Password string

	// Nickname contains the display name.
	Nickname string

	// Nonce is a per-user secret for signing JWT.
	Nonce []byte
}
