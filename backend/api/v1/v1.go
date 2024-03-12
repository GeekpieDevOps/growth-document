package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Mount(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/login", Login(db))
	r.POST("/register", Register(db))
	r.PUT("/update", Update(db))
	r.DELETE("/delete", Delete(db))
}
