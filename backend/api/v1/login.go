package v1

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}

func Login(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to read request body"})
			return
		}

		var req LoginRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to parse request body"})
			return
		}

		var user models.User
		result := db.Where("email = ? AND password = ?", req.Email, req.Password).First(&user)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.String(http.StatusNotFound, "")
			return
		}

		response := LoginResponse{
			Code:    200,
			Message: "登录成功",
		}
		response.Data.Type = "Student"
		response.Data.ID = "123456"

		c.JSON(http.StatusOK, response)
	}
}
