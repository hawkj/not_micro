package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/hawkj/not_micro/config"
)

func NewRedis(c config.Redis) *redis.Client {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: "",
		DB:       0,
	})
	return rdb
}
