package adapter

import (
	"context"
	"event_sourcing_payment/constant"
	"event_sourcing_payment/infrastructure/redis_client"
	"event_sourcing_payment/package/locker"
	"event_sourcing_payment/package/logger"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type IAdapter interface {
	RedisClient() *redis.Client
	Locker() locker.Locker
}

type Adapter struct {
	redisClient *redis.Client
	locker      locker.Locker
}

func NewAdapter(ctx context.Context, cfg *constant.Config) (IAdapter, error) {
	log := logger.FromContext(ctx)
	redisClient, err := redis_client.NewRedisClient(ctx, &cfg.Redis)
	if err != nil {
		log.Error("Failed to initialize Redis client", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize Redis client: %w", err)
	}

	locker := locker.NewLocker(ctx, redisClient)

	return &Adapter{
		redisClient: redisClient,
		locker:      locker,
	}, nil
}

func (a *Adapter) RedisClient() *redis.Client {
	return a.redisClient
}

func (a *Adapter) Locker() locker.Locker {
	return a.locker
}
