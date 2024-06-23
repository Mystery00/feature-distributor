package redis

import (
	"context"
	"errors"
	"feature-distributor/common/env"
	"fmt"
	v9 "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var rdb *v9.Client

func InitRedis() {
	address := viper.GetString(env.RedisAddress)
	password := viper.GetString(env.RedisPassword)

	rdb = v9.NewClient(&v9.Options{
		Addr:     address,
		Password: password,
	})
}

func Set(ctx context.Context, key string, value string, expiration time.Duration) (err error) {
	prefix := viper.GetString(env.RedisPrefix)
	redisKey := fmt.Sprintf("%s%s", prefix, key)
	err = rdb.Set(ctx, redisKey, value, expiration).Err()
	return
}

func Get(ctx context.Context, key string) (value *string, err error) {
	prefix := viper.GetString(env.RedisPrefix)
	redisKey := fmt.Sprintf("%s%s", prefix, key)
	bytes, err := rdb.Get(ctx, redisKey).Bytes()
	if err != nil && !errors.Is(err, v9.Nil) {
		return nil, err
	}
	if bytes == nil {
		return nil, nil
	}
	val := string(bytes)
	value = &val
	return
}

func Expire(ctx context.Context, key string, expiration time.Duration) (err error) {
	prefix := viper.GetString(env.RedisPrefix)
	redisKey := fmt.Sprintf("%s%s", prefix, key)
	err = rdb.Expire(ctx, redisKey, expiration).Err()
	return
}

func Del(ctx context.Context, key string) (err error) {
	prefix := viper.GetString(env.RedisPrefix)
	redisKey := fmt.Sprintf("%s%s", prefix, key)
	err = rdb.Del(ctx, redisKey).Err()
	return
}
