package database

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	config *config.Config
}

func New(config *config.Config) *DB {
	return &DB{
		config: config,
	}
}

func (d *DB) InitializeDB() (*gorm.DB, error) {
	db, err := d.connect()
	if err != nil {
		return nil, err
	}

	return db, migrate(db)
}

func (d *DB) connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		d.config.CardeaDB.Host,
		d.config.Secrets.CardeaDBCredentials.Username,
		d.config.Secrets.CardeaDBCredentials.Password,
		d.config.CardeaDB.Database,
		d.config.CardeaDB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("failed to connect to cardea db", zap.Error(err))
	}

	return db, err
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate()
}
