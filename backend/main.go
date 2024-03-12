package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/api"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DefaultResponse struct {
	Code    int64                  `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

type RequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	ID       string `json:"id"`
}

type UpdateResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       string `json:"id"`
}

type User struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}

func UpdateUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var user User
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
		ID:       user.Data.ID,
	}
	response.Message = "User updated successfully"

	c.JSON(http.StatusOK, response)
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	group := r.Group("/api")

	api.Mount(group, db)

	r.PUT("/api/v1/update", func(c *gin.Context) {
		UpdateUser(c, db)
	})

	return r
}

func main() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai" //postgres 是 PostgreSQL 的默认用户名和数据库名称
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})                                                                     //"   "应该填链接，由于我不知道postgres的链接形式，所以这里只是一个示例
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")

	db.AutoMigrate(&User{})

	r := setupRouter(db)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
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
