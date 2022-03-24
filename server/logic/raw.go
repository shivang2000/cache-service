package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func SetKey(rbd *redis.Client, ctx context.Context, key string, value string) error {
	// value = "shivang:" + value
	err := rbd.Set(ctx, key, "shivang:"+value, 0).Err()
	return err
}

func GetValue(rbd *redis.Client, ctx context.Context, key string) (string, error) {
	val, err := rbd.Get(ctx, key).Result()
	return val, err
}
