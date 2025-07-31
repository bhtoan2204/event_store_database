package projection

import (
	"context"
	"event_sourcing_payment/constant"
	"event_sourcing_payment/package/logger"
	"fmt"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ProjectionConnection struct {
	db *gorm.DB
}

func NewProjectionConnection(ctx context.Context, cfg *constant.PostgresConfig) (*ProjectionConnection, error) {
	log := logger.FromContext(ctx)
	if cfg == nil {
		log.Fatal("PostgresConfig is nil")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect to Postgres", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error("failed to get sql.DB from gorm DB", zap.Error(err))
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Error("failed to use otelgorm plugin", zap.Error(err))
	}

	log.Info("Successfully connected to Postgres database")

	return &ProjectionConnection{
		db: db,
	}, nil
}

func (p *ProjectionConnection) SyncTable(models ...interface{}) error {
	return p.db.AutoMigrate(models...)
}

func (p *ProjectionConnection) GetDB() *gorm.DB {
	return p.db
}
