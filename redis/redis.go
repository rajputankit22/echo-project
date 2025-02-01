package redis

import (
	"context"
	"echo-project/config"
	"echo-project/logger"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx                 = context.Background()
	newRedisCacheClient = newRedisCache
)

// CacheInterface - interface for cache
type CacheInterface interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, expiration int) error
}

type redisCache struct {
	client *redis.Client
}

// Constructor function for redisCache
func NewRedisCacheAdapter() (CacheInterface, error) {
	return &redisCache{
		client: newRedisCacheClient(),
	}, nil
}

// Get - get value from cache
func (r *redisCache) Get(key string) ([]byte, error) {
	val, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// Set - set value in cache
func (r *redisCache) Set(key string, value []byte, expiration int) error {
	return r.client.Set(ctx, key, value, time.Duration(expiration*int(time.Second))).Err()
}

// newRedisCache - create a new redis client
func newRedisCache() *redis.Client {
	opts := &redis.Options{
		Addr:     config.Config().Cache.Host + ":" + config.Config().Cache.Port,
		Password: config.Config().Cache.Password,
		DB:       config.Config().Cache.DB,
	}
	client := redis.NewClient(opts)

	// Ping the Redis server to check the connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Error("", "Failed to connect to Redis:", err)
	} else {
		logger.Trace("Successfully connected to Redis")
	}

	return client
}
