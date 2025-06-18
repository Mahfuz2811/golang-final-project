package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisCache) Set(key string, value string, expiration time.Duration) {
	r.client.Set(key, value, expiration)
}
