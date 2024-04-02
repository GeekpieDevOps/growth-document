package v1

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string
	Password string
	Nickname string
}

func Register(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to read request body"})
			return
		}

		var req RegisterRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to parse request body"})
			return
		}

		var user models.User
		result := db.Where("email = ?", req.Email).First(&user)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.String(http.StatusConflict, "")
			return
		}

		uuid, err := uuid.NewRandom()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": ""})
			return
		}

		user = models.User{
			UUID:     uuid,
			Email:    req.Email,
			Password: req.Password,
			Nickname: req.Nickname,
		}
		result = db.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": ""})
	}
}
