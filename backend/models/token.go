package models

import (
	"github.com/google/uuid"
)

// A Token is a JWT.
type Token struct {
	// ID is the unique identifier of the token.
	ID uuid.UUID `gorm:"primaryKey"`

	// UUID is the unique identifier of the owner.
	UUID uuid.UUID

	// Nonce is used to sign and parse the token.
	Nonce []byte

	// Token contains signed JWT.
	Token string
}
