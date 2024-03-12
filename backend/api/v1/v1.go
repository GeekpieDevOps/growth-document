package v1

import (
	"github.com/GeekpieDevOps/growth-document/backend/api/v1/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Mount(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/login", Login(db))
	r.POST("/register", Register(db))
	r.PUT("/update", Update(db))
	r.DELETE("/delete", Delete(db))

	user.Mount(r.Group("/user"), db)
}
