// Package user contains implementaions for the /api/v1/user API family.
package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Mount mounts handlers under r.
// It also configures them to use db for database access.
func Mount(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/sign_up", SignUp(db))
	r.POST("/sign_in", SignIn(db))
}
