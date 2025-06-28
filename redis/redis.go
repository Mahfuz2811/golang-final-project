package redis

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func NewRedisClient() (*redis.Client, error) {
	host := getEnv("REDIS_HOST", "localhost")
	port := getEnv("REDIS_PORT", "6379")
	password := getEnv("REDIS_PASSWORD", "redispassword")
	dbStr := getEnv("REDIS_DB", "0")

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		db = 0
	}

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return client, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
