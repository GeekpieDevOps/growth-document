package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Mount(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/sign_up", SignUp(db))
	r.POST("/sign_in", SignIn(db))
}
