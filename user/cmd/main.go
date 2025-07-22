package main

import (
	"context"
	"event_sourcing_user/config"
	"event_sourcing_user/presentation"
	"log"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Printf("application went wrong. Panic err=%v", r)
		}
	}()

	err := run(ctx)
	done()
	if err != nil {
		log.Printf("realMain has failed with err=%v", err)
		return
	}
}

func run(ctx context.Context) error {
	godotenv.Load(".env")
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}

	app, err := presentation.NewApp(cfg)
	if err != nil {
		return err
	}

	return app.Start(ctx)
}
