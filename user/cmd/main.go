package main

import (
	"context"
	"event_sourcing_user/config"
	"event_sourcing_user/package/logger"
	"event_sourcing_user/presentation"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	log := logger.FromContext(ctx)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Error("application went wrong. Panic err=%v", zap.Any("err", r))
		}
	}()

	err := run(ctx)
	done()
	if err != nil {
		log.Error("realMain has failed with err=%v", zap.Error(err))
		return
	}
}

func run(ctx context.Context) error {
	log := logger.FromContext(ctx)
	godotenv.Load(".env")
	cfg, err := config.InitConfig()
	if err != nil {
		log.Error("Error initializing config", zap.Error(err))
		return err
	}

	app, err := presentation.NewApp(ctx, cfg)
	if err != nil {
		log.Error("Error initializing app", zap.Error(err))
		return err
	}

	return app.Start(ctx)
}
