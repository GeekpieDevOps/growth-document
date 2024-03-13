// Package models contains database models.
package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) (err error) {
	return db.AutoMigrate(&User{})
}
