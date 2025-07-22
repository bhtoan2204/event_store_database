package persistent

import (
	"event_sourcing_user/constant"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersistentConnection struct {
	db *gorm.DB
}

func NewPersistentConnection(cfg *constant.PostgresConfig) (*PersistentConnection, error) {
	if cfg == nil {
		log.Fatal("PostgresConfig is nil")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatalf("failed to use otelgorm plugin: %v", err)
	}

	log.Println("Successfully connected to Postgres database")

	return &PersistentConnection{
		db: db,
	}, nil
}

func (p *PersistentConnection) SyncTable(models ...interface{}) error {
	return p.db.AutoMigrate(models...)
}

func (p *PersistentConnection) GetDB() *gorm.DB {
	return p.db
}
