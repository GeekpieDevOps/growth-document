package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/GeekpieDevOps/growth-document/backend/api"
	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	slogGorm "github.com/orandin/slog-gorm"
	slogGin "github.com/samber/slog-gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRouter(logger *slog.Logger, db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(slogGin.New(logger))
	r.Use(gin.Recovery())

	api.Mount(r.Group("/api"), db)

	return r
}

func main() {
	logger := slog.Default()

	dsn := os.Getenv("GD_DSN")
	if dsn == "" {
		logger.Error("environment variable DSN is not set")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: slogGorm.New(
			slogGorm.WithLogger(logger),
		),
	})

	if err != nil {
		// error already logged above
		os.Exit(1)
	}

	db.AutoMigrate(&models.User{})

	r := setupRouter(logger, db)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
