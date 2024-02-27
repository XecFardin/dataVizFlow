package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis address
		Password: "",               // Redis password
		DB:       0,                // Redis database
	})
}

func CacheData(key string, value interface{}) error {
	ctx := context.Background()
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetCachedData(key string) (string, error) {
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}
	return val, nil
}
