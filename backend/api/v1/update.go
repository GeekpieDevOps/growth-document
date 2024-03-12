package v1

import (
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       string `json:"id"`
}

func Update(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user models.User
		result := db.First(&user, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"code": 500, "message": "User not found"})
			return
		}

		var requestData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			Nickname string `json:"nickname"`
		}
		err := c.ShouldBindJSON(&requestData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 500, "message": "Invalid request data"})
			return
		}

		result = db.Save(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to update user"})
			return
		}

		response := UpdateResponse{
			Code:     200,
			Email:    requestData.Email,
			Password: requestData.Password,
			ID:       user.ID,
		}
		response.Message = "User updated successfully"

		c.JSON(http.StatusOK, response)
	}
}
