package initializer

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}
}
