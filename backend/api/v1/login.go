package v1

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
			// 处理读取请求体错误
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to read request body"})
			return
		}

		fmt.Print(string(body))

		response := LoginResponse{
			Code:    200,
			Message: "登录成功",
		}
		response.Data.Type = "Student"
		response.Data.ID = "123456"

		c.JSON(http.StatusOK, response)
	}
}
