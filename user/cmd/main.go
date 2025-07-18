package main

import (
	"event_sourcing_user/config"
	"event_sourcing_user/infrastructure/persistent"
	"event_sourcing_user/infrastructure/persistent/persistent_object"
	"event_sourcing_user/infrastructure/persistent/repository"

	"github.com/joho/godotenv"
)

func main() {
	run()
}

func run() {
	godotenv.Load(".env")
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	persistentConnection := persistent.NewPersistentConnection(&cfg.Postgres)
	persistentConnection.SyncTable(&persistent_object.User{})

	repositoryFactory := repository.NewRepositoryFactory(persistentConnection)
	if err != nil {
		panic(err)
	}

}
