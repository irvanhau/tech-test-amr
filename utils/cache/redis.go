package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type redisClient struct {
	rdb *redis.Client
}

type Redis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
}

func InitRedis() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &redisClient{rdb: client}
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.rdb.Set(ctx, key, jsonData, 10*time.Minute).Err()
	return err
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
