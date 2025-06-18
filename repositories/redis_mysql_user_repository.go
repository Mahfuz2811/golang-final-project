package repositories

import (
	"encoding/json"
	"final-golang-project/models"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisMySQLUserRepository struct {
	repo  UserRepository
	redis *redis.Client
}

func NewRedisMySQLUserRepository(repo UserRepository, redisClient *redis.Client) *RedisMySQLUserRepository {
	return &RedisMySQLUserRepository{
		repo:  repo,
		redis: redisClient,
	}
}

func (r *RedisMySQLUserRepository) Create(user models.User) error {
	fmt.Println("Creating")
	return r.repo.Create(user)
}

func (r *RedisMySQLUserRepository) GetByEmail(email string) (*models.User, error) {
	cacheKey := "user:" + email // user:test@gmail.com = {}
	value, err := r.redis.Get(cacheKey).Result()
	if err == nil {
		var cachedUser models.User
		// string data has to be mapped with user model
		if err := json.Unmarshal([]byte(value), &cachedUser); err == nil {
			return &cachedUser, nil
		}
	}

	user, error := r.repo.GetByEmail(email)
	if error != nil || user == nil {
		return nil, error
	}

	fmt.Println("Set cache")
	// cache user for future
	userJson, _ := json.Marshal(user)
	r.redis.Set(cacheKey, userJson, 10*time.Minute)

	return user, nil
}
