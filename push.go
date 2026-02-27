package redisctl

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

func RPushWithDefaultTimeout(client redis.Cmdable, key string, values ...interface{}) error {
	if client == nil {
		return ClientIsNil
	}

	if key == "" {
		return KeyIsEmpty
	}

	if len(values) == 0 {
		return errors.New("RPush values cannot be empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	_, err := client.RPush(ctx, key, values).Result()

	return err
}
