// Package main contains the program entry point.
//
// Behaviors of the program can be configured
// by the following environment variables:
//   - GD_DRV: database engine to use. Valid values are
//     sqlite for SQLite and postgres for PostgreSQL.
//   - GD_DSN: specifies which database to use. The exact format depends on the
//     database engine used.
package main

import (
	"errors"
	"log/slog"
	"os"
	"strings"

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
	db, err := openDatabase()
	if err != nil {
		// errors are already logged
		os.Exit(1)
	}

	if err := models.AutoMigrate(db); err != nil {
		// errors are already logged
		os.Exit(1)
	}

	if err := setupRouter(db).Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func openDatabase() (db *gorm.DB, err error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	dsn := os.Getenv("GD_DSN")
	drv := os.Getenv("GD_DRV")
	if databaseUrl != "" {
		dsn = databaseUrl
		drv = dsn[:strings.Index(dsn, ":")]
	}
	if dsn == "" {
		err = errors.New("environment variable GD_DSN is empty or not set. Or DATABASE_URL is not properly set")
		slog.Error(err.Error())
		return
	}

	var dialector gorm.Dialector
	switch drv {
	case "":
		slog.Warn("environment variable GD_DRV is empty or not set, defaulting to SQLite. Or DATABASE_URL is not properly set")
		fallthrough
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		err = errors.New("environment variable GD_DRV has invalid value, should be one of `sqlite' or `postgres'. Or DATABASE_URL is not properly set")
		slog.Error(err.Error())
		return
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: slogGorm.New(),
	})

	return
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(slogGin.New(slog.Default()))
	r.Use(gin.Recovery())

	api.Mount(r.Group("/api"), db)

	return r
}
