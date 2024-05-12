package database

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
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

func (d *DB) Initialize() *gorm.DB {
	db, err := d.connect()
	if err != nil {
		zap.L().Fatal("failed to connect to cardea db", zap.Error(err))
	}

	err = migrate(db)
	if err != nil {
		zap.L().Fatal("failed to migrate to cardea db", zap.Error(err))
	}

	zap.L().Info("database initialized successfully")

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

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Workout{},
		&entity.Exercise{},
		&entity.Profile{},
	)
}
