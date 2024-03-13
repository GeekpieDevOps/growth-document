package models

import (
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `gorm:"primaryKey"`
	ID       string
	Email    string
	Password string
	Nickname string
	Nonce    []byte
}
