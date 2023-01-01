package gredis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var RedisConn redis.UniversalClient
var ctx = context.Context(context.Background())

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      []string{setting.RedisSetting.Host},
		PoolSize:   setting.RedisSetting.PoolSize,
		MaxRetries: setting.RedisSetting.MaxRetries,
		Password:   setting.RedisSetting.Password,
	})
	_, err := RedisConn.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// Set a key/value
func Set(key string, data interface{}, time time.Duration) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	statusCmd := RedisConn.Set(ctx, key, value, time)
	_, err = statusCmd.Result()
	if err != nil {
		return err
	}
	return nil
}

// Exists check a key
func Exists(key string) bool {
	intCmd := RedisConn.Exists(ctx, key)
	_, err := intCmd.Result()
	if err != nil {
		return false
	}
	return true
}

// Get get a key
func Get(key string) ([]byte, error) {
	stringCmd := RedisConn.Get(ctx, key)
	reply, err := stringCmd.Bytes()
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	intCmd := RedisConn.Del(ctx, key)
	_, err := intCmd.Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
