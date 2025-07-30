package redis_client

import (
	"context"
	"event_sourcing_gateway/package/settings"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *settings.RedisConfig) (*redis.Client, error) {
	var (
		redisClient *redis.Client
		err         error
	)

	redisClient, err = newStandAlone(cfg)
	if err != nil {
		return nil, err
	}

	cmd := redisClient.Ping(context.Background())
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	return redisClient, nil
}

func newStandAlone(cfg *settings.RedisConfig) (*redis.Client, error) {
	opts, err := redis.ParseURL(cfg.ConnectionURL)
	if err != nil {
		return nil, fmt.Errorf("parseURl failed err=%w", err)
	}

	opts.PoolSize = cfg.PoolSize
	opts.DialTimeout = time.Duration(cfg.DialTimeoutSeconds) * time.Second
	opts.ReadTimeout = time.Duration(cfg.ReadTimeoutSeconds) * time.Second
	opts.WriteTimeout = time.Duration(cfg.WriteTimeoutSeconds) * time.Second
	opts.ConnMaxIdleTime = time.Duration(cfg.IdleTimeoutSeconds) * time.Second
	opts.MaxIdleConns = cfg.MaxIdleConn
	opts.MaxActiveConns = cfg.MaxActiveConn

	return redis.NewClient(opts), nil
}
