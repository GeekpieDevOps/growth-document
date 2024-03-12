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

func main() {
	dsn := os.Getenv("GD_DSN")
	if dsn == "" {
		slog.Error("environment variable GD_DSN is not set")
		os.Exit(1)
	}

	db, err := setupPostgres(dsn)
	if err != nil {
		// errors are already logged
		os.Exit(1)
	}

	if err := setupRouter(db).Run(); err != nil {
		log.Fatal(err)
	}
}

func setupPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: slogGorm.New(),
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return db, nil
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(slogGin.New(slog.Default()))
	r.Use(gin.Recovery())

	api.Mount(r.Group("/api"), db)

	return r
}
