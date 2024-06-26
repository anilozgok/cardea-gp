package database

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/pkg/reader"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"strings"
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

	err = injectInitData(db)
	if err != nil {
		zap.L().Fatal("failed to inject initial data to cardea db", zap.Error(err))
	}

	err = injectInitFoodData(db)
	if err != nil {
		zap.L().Fatal("failed to inject initial food data to cardea db", zap.Error(err))
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
		&entity.Photo{},
		&entity.Diet{},
		&entity.Meal{},
		&entity.Food{},
	)
}

func injectInitData(db *gorm.DB) error {
	rows, err := reader.CSV("exercise.csv")
	if err != nil {
		return err
	}

	exercises := make([]entity.Exercise, 0)

	for i, r := range rows {
		if i == 0 {
			continue
		}

		convertedGiftURL := convertGIFURL(r[2])

		exercise := entity.Exercise{
			Name:      r[3],
			BodyPart:  r[0],
			Target:    r[4],
			Equipment: r[1],
			Gif:       convertedGiftURL,
		}

		exercises = append(exercises, exercise)
	}

	tx := db.Exec("TRUNCATE ONLY exercises RESTART IDENTITY;")
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.CreateInBatches(&exercises, 100)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func injectInitFoodData(db *gorm.DB) error {
	rows, err := reader.CSV("diet.csv")
	if err != nil {
		return err
	}

	foods := make([]entity.Food, 0)

	for i, r := range rows {
		if i == 0 {
			continue
		}

		avgServingSize, err := strconv.ParseFloat(r[1], 64)
		if err != nil {
			return fmt.Errorf("error parsing avg serving size: %v", err)
		}

		calories, err := strconv.ParseFloat(r[2], 64)
		if err != nil {
			return fmt.Errorf("error parsing calories: %v", err)
		}

		food := entity.Food{
			Name:           r[0],
			AvgServingSize: avgServingSize,
			Calories:       calories,
			Category:       r[3],
		}

		foods = append(foods, food)
	}

	tx := db.Exec("TRUNCATE ONLY foods RESTART IDENTITY;")
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.CreateInBatches(&foods, 100)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func convertGIFURL(oldURL string) string {
	splits := strings.Split(oldURL, "/")
	return fmt.Sprintf("https://giphy.com/embed/%s", splits[len(splits)-2])
}
