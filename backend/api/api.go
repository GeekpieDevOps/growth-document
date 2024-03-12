package api

import (
	v1 "github.com/GeekpieDevOps/growth-document/backend/api/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Mount(r *gin.RouterGroup, db *gorm.DB) {
	v1.Mount(r.Group("/v1"), db)
}
