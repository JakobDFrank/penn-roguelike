package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	_redisAddr = "redis:6379"
)

// redisCache is an implementation of Cacher that uses Redis
type redisCache struct {
	client *redis.Client
}

func (r *redisCache) GetBytes(ctx context.Context, key string) ([]byte, error) {
	cmd := r.client.Get(ctx, key)

	return cmd.Bytes()
}

func (r *redisCache) SetBytes(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// NewRedisCache creates a new instance of Cacher that uses Redis
func NewRedisCache() (Cacher, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     _redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &redisCache{rdb}, nil
}

var _ Cacher = (*redisCache)(nil)
