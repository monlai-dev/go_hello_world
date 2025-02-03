package cache

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"os"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return RedisClient
}
