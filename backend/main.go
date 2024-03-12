package main

import (
	"errors"
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
	db, err := setupPostgres()
	if err != nil {
		// errors are already logged
		os.Exit(1)
	}

	if err := setupRouter(db).Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func setupPostgres() (*gorm.DB, error) {
	dsn := os.Getenv("GD_DSN")
	if dsn == "" {
		err := errors.New("environment variable GD_DSN is empty or not set")
		slog.Error(err.Error())
		return nil, err
	}

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
