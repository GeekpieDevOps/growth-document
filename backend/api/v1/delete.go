package v1

import (
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user models.User
		result := db.First(&user, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"code": 500, "message": "User not found"})
			return
		}

		result = db.Delete(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "User deleted successfully"})
	}
}
