package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RateLimitMiddleware struct {
	redisClient *redis.Client
	limit       int
	window      time.Duration
}

func NewRateLimitMiddleware(redisClient *redis.Client, limit int, window time.Duration) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		redisClient: redisClient,
		limit:       limit,
		window:      window,
	}
}

func (r *RateLimitMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		identifier := c.ClientIP()
		key := "rate_limit:" + identifier

		now := time.Now().UnixMilli()
		windowStart := now - r.window.Milliseconds()

		pipe := r.redisClient.Pipeline()

		pipe.ZRemRangeByScore(ctx, key, "0", strconv.FormatInt(windowStart, 10))

		pipe.ZAdd(ctx, key, redis.Z{
			Score:  float64(now),
			Member: now,
		})

		countCmd := pipe.ZCard(ctx, key)

		pipe.Expire(ctx, key, r.window*2)

		_, err := pipe.Exec(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		count := countCmd.Val()
		if int(count) > r.limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}
