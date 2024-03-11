package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UUID     uuid.UUID
	Email    string
	Password string
	Nickname string
}
