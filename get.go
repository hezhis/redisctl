package redisctl

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisValue struct {
	val string
	err error
}

func NewRedisValue(cmd *redis.StringCmd) RedisValue {
	val, err := cmd.Result()
	if errors.Is(err, redis.Nil) {
		val = ""
		err = nil
	}
	return RedisValue{
		val: val,
		err: err,
	}
}

func (r *RedisValue) Int32() (int32, error) {
	if r.err != nil {
		return 0, r.err
	}

	if r.val == "" {
		return 0, nil
	}

	v, err := strconv.ParseInt(r.val, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), err
}

func (r *RedisValue) Uint32() (uint32, error) {
	if r.err != nil {
		return 0, r.err
	}

	if r.val == "" {
		return 0, nil
	}

	v, err := strconv.ParseUint(r.val, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func (r *RedisValue) Int64() (int64, error) {
	if r.err != nil {
		return 0, r.err
	}
	if r.val == "" {
		return 0, nil
	}
	return strconv.ParseInt(r.val, 10, 64)
}

func (r *RedisValue) Uint64() (uint64, error) {
	if r.err != nil {
		return 0, r.err
	}
	if r.val == "" {
		return 0, nil
	}

	return strconv.ParseUint(r.val, 10, 64)
}

func (r *RedisValue) String() (string, error) {
	return r.val, r.err
}

func (r *RedisValue) Bool() (bool, error) {
	if r.err != nil {
		return false, r.err
	}
	if r.val == "" {
		return false, nil
	}
	return strconv.ParseBool(r.val)
}

func (r *RedisValue) Bytes() ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}
	if r.val == "" {
		return nil, nil
	}
	return []byte(r.val), nil
}

func (r *RedisValue) Error() error {
	return r.err
}
