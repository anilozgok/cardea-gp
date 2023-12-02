package main

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	l, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("fail to build log. err: %s", err))
	}

	zap.ReplaceGlobals(l)
}

func main() {
	appServer := server.NewAppServer()
	appServer.Start()

	configs, err := config.Get()
	if err != nil {
		zap.L().Fatal("failed to read configs", zap.Error(err))
	}

	db := connectDB(configs, err)

	gracefulShutdown(appServer)
}

func connectDB(config *config.Config, err error) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		config.CardeaDB.Host,
		config.Secrets.CardeaDBCredentials.Username,
		config.Secrets.CardeaDBCredentials.Password,
		config.CardeaDB.Database,
		config.CardeaDB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("failed to connect to cardea db", zap.Error(err))
	}

	err = db.AutoMigrate()
	if err != nil {
		zap.L().Fatal("failed to migrate to cardea db", zap.Error(err))
	}

	return db
}

func gracefulShutdown(appServer *server.AppServer) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-signalCh

	zap.L().Info("shutting down server")
	if err := appServer.Shutdown(); err != nil {
		zap.L().Error("error occurred while shutting down server", zap.Error(err))
	}
}
