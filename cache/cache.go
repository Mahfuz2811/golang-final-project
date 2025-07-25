package cache

import "time"

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration)
}
