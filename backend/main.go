package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/api"
	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	group := r.Group("/api")

	api.Mount(group, db)

	return r
}

func main() {
	// FIXME: default user name and database name is used here
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")

	db.AutoMigrate(&models.User{})

	r := setupRouter(db)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
