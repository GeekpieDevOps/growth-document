package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/GeekpieDevOps/growth-document/backend/api"
	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	slogGorm "github.com/orandin/slog-gorm"
	slogGin "github.com/samber/slog-gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := setupDatabase()
	if err != nil {
		// errors are already logged
		os.Exit(1)
	}

	if err := setupRouter(db).Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func setupDatabase() (db *gorm.DB, err error) {
	dsn := os.Getenv("GD_DSN")
	if dsn == "" {
		err = errors.New("environment variable GD_DSN is empty or not set")
		slog.Error(err.Error())
		return
	}

	var dialector gorm.Dialector
	switch os.Getenv("GD_DRV") {
	case "":
		slog.Warn("environment variable GD_DRV is empty or not set, defaulting to SQLite")
		fallthrough
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		err = errors.New("environment variable GD_DRV has invalid value, should be one of `sqlite' or `postgres'")
		slog.Error(err.Error())
		return
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: slogGorm.New(),
	})
	if err != nil {
		return
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(slogGin.New(slog.Default()))
	r.Use(gin.Recovery())

	api.Mount(r.Group("/api"), db)

	return r
}
