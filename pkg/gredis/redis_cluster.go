package gredis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func RedisCluster(key string, value int) error {
	// Create a new lock client.
	keyValue := RedisConn.Get(ctx, key)
	keyInt, _ := keyValue.Int()
	// 需要给倒计时自动释放掉，但是有问题。最好使用lua脚本，进行删除，因为如果进程挂了，就GG了。
	set, err := RedisConn.SetNX(ctx, key, strconv.Itoa(keyInt+value), 1000*time.Second).Result()
	if err != nil {
		return err
	}
	vals, err := RedisConn.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  2,
	}).Result()
	if err != nil {
		return err
	}
	fmt.Printf("set value %v \n", set)
	fmt.Printf("vals value %v \n", vals)
	return nil
}
