package locker

import (
	"context"
	"event_sourcing_payment/package/logger"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Locker interface {
	AcquireLock(ctx context.Context, key string, value string, expiration time.Duration, retryDelay time.Duration, timeout time.Duration) (bool, error)
	ReleaseLock(ctx context.Context, key string, value string) (bool, error)
}

type locker struct {
	client *redis.Client
}

func NewLocker(ctx context.Context, client *redis.Client) Locker {
	log := logger.FromContext(ctx)
	log.Info("Initializing locker")

	return &locker{
		client: client,
	}
}

func (l *locker) AcquireLock(ctx context.Context, key string, value string, expiration time.Duration, retryDelay time.Duration, timeout time.Duration) (bool, error) {
	log := logger.FromContext(ctx)
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	success := false
	var err error

	defer func() {
		if err != nil || !success {
			log.Error("AcquireLock got error",
				zap.Error(err), zap.String("key", key), zap.Bool("can", success))
		}
	}()
	for {
		success, err = l.client.SetNX(ctx, key, value, expiration).Result()
		if err != nil {
			return false, err
		}
		if success {
			return true, nil
		}

		select {
		case <-time.After(retryDelay):

		case <-timer.C:
			err = fmt.Errorf("timeout exceeded while trying to acquire lock")
			return false, err
		case <-ctx.Done():
			err = ctx.Err()
			return false, err
		}
	}
}

func (l *locker) ReleaseLock(ctx context.Context, key string, value string) (bool, error) {
	log := logger.FromContext(ctx)

	script := redis.NewScript(`
        if redis.call("get", KEYS[1]) == ARGV[1] then
            return redis.call("del", KEYS[1])
        else
            return 0
        end
    `)
	result, err := script.Run(ctx, l.client, []string{key}, value).Result()
	if err != nil {
		log.Error("ReleaseLock key %s val %s error %w", zap.Error(err))
		return false, fmt.Errorf("ReleaseLock key %s val %s error %w", key, value, err)
	}

	hit := result.(int64)
	if !(hit > 0) {
		log.Error("ReleaseLock key %s val %s hit %d", zap.Int64("hit", hit))
		return false, fmt.Errorf("ReleaseLock key %s val %s hit %d", key, value, hit)
	}

	return hit > 0, nil
}
