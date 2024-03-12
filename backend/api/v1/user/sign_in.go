package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignIn(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}
