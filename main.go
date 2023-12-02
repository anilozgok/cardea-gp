package main

import (
	"fmt"
	"github.com/anilozgok/cardea-gp/internal/config"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	configs, err := config.Get()
	if err != nil {
		zap.L().Fatal("failed to read configs", zap.Error(err))
	}

	cDB := database.New(configs)

	db, err := cDB.Initialize()
	if err != nil {
		zap.L().Fatal("failed to initialize database", zap.Error(err))
	}

	appServer := server.NewAppServer(db)
	appServer.Start()

	gracefulShutdown(appServer)
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
