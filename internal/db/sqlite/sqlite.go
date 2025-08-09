package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"guitar-go/internal/config"
)

type SQLiteDB struct {
	config *config.SQLiteConfig
	db     *gorm.DB
}

func NewSQLiteDB(cfg *config.SQLiteConfig) *SQLiteDB {
	return &SQLiteDB{config: cfg}
}

func (s *SQLiteDB) Connect() error {
	db, err := gorm.Open(sqlite.Open(s.config.Path), &gorm.Config{})
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *SQLiteDB) GetDB() *gorm.DB {
	return s.db
}

func (s *SQLiteDB) AutoMigrate(models ...interface{}) error {
	return s.db.AutoMigrate(models...)
}

func (p *SQLiteDB) Ping() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
