package cache

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"os"
)

var RedisClient *redis.Client
var ctx = context.Background()

var redisKey string

func init() {
	if os.Getenv("ENV") == "staging" {
		redisKey = "RENDER_REDIS_URL"
		return
	}

	redisKey = "REDIS_URL"

}

func ConnectRedis() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisKey,
		Password: "",
		DB:       0,
	})

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return RedisClient
}
