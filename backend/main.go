package main

import (
	"log"
	"os"

	"github.com/GeekpieDevOps/growth-document/backend/api"
	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	api.Mount(r.Group("/api"), db)

	return r
}

func main() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("ERROR: Couldn't connect to database: environment variable DSN is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{})

	r := setupRouter(db)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
