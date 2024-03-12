package api

import (
	"net/http"

	v1 "github.com/GeekpieDevOps/growth-document/backend/api/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Mount(r *gin.RouterGroup, db *gorm.DB) {
	r.Use(corsMiddleware())

	r.POST("/", Hello(db))

	v1.Mount(r.Group("/v1"), db)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
