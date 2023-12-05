package database

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/entities"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type DB struct {
	config *config.Config
}

func New(config *config.Config) *DB {
	return &DB{
		config: config,
	}
}

func (d *DB) Initialize() *gorm.DB {
	time.Sleep(5 * time.Second) // wait for db to be ready
	db, err := d.connect()
	if err != nil {
		zap.L().Fatal("failed to connect to cardea db", zap.Error(err))
	}

	err = migrate(db)
	if err != nil {
		zap.L().Fatal("failed to migrate to cardea db", zap.Error(err))
	}

	return db
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
	return db.AutoMigrate(
		&entities.User{},
	)
}
