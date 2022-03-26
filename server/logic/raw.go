package logic

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const ttlDuration time.Duration = 300000000000 // 300s

func SetKey(rbd *redis.Client, ctx context.Context, key string, value string) error {
	value = "shivang:" + value

	err := rbd.Set(ctx, key, value, ttlDuration).Err()
	return err
}

func GetValue(rbd *redis.Client, ctx context.Context, key string) (string, error) {
	val, err := rbd.Get(ctx, key).Result()
	return val, err
}
