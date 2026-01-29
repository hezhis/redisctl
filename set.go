package redisctl

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

const DefaultTimeout = time.Second * 3

var (
	ClientIsNil = errors.New("redis client is nil")
	KeyIsEmpty  = errors.New("key is empty")
)

func Set(client *redis.Client, key string, value interface{}, expire, timeout time.Duration) error {
	if client == nil {
		return ClientIsNil
	}

	if key == "" {
		return KeyIsEmpty
	}

	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.Set(ctx, key, value, expire).Result()

	return err
}

func HSet(client *redis.Client, key string, field string, value interface{}, timeout time.Duration) error {
	if client == nil {
		return ClientIsNil
	}
	if key == "" {
		return KeyIsEmpty
	}
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.HSet(ctx, key, field, value).Result()
	return err
}
