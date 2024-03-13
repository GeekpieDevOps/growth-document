// Package models contains database models.
package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) (err error) {
	if err = db.AutoMigrate(&Token{}); err != nil {
		return
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		return
	}

	return
}
