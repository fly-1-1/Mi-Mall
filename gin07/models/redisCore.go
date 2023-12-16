package models

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var redisCoretxt = context.Background()
var (
	RedisDb *redis.Client
)

func init() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisDb.Ping(redisCoretxt).Result()
	if err != nil {
		println(err)
	}
}
