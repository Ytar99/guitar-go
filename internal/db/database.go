package db

import (
	"fmt"

	"gorm.io/gorm"

	"guitar-go/internal/config"
	"guitar-go/internal/db/postgres"
	"guitar-go/internal/db/sqlite"
)

type Database interface {
	Connect() error
	GetDB() *gorm.DB
	AutoMigrate(models ...interface{}) error
	Ping() error
}

func NewDatabase(cfg *config.Config) (Database, error) {
	switch cfg.Database.Driver {
	case "sqlite":
		return sqlite.NewSQLiteDB(&cfg.Database.SQLite), nil
	case "postgres":
		return postgres.NewPostgresDB(&cfg.Database.Postgres), nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}
}
