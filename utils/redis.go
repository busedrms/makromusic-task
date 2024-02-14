package utils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func InitializeRedis(address, password string, db int) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	return err
}

func SetToken(key, token string, expiration time.Duration) error {
	return rdb.Set(context.Background(), key, token, expiration).Err()
}

func GetToken(key string) (string, error) {
	return rdb.Get(context.Background(), key).Result()
}

func DeleteToken(key string) error {
	return rdb.Del(context.Background(), key).Err()
}

func CloseRedis() {
	if rdb != nil {
		rdb.Close()
	}
}
