package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"guitar-go/internal/config"
)

type PostgresDB struct {
	config *config.PostgresConfig
	db     *gorm.DB
}

func NewPostgresDB(cfg *config.PostgresConfig) *PostgresDB {
	return &PostgresDB{config: cfg}
}

func (p *PostgresDB) Connect() error {
	db, err := gorm.Open(postgres.Open(p.config.URL), &gorm.Config{})
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *PostgresDB) GetDB() *gorm.DB {
	return p.db
}

func (p *PostgresDB) AutoMigrate(models ...interface{}) error {
	return p.db.AutoMigrate(models...)
}

func (p *PostgresDB) Ping() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
