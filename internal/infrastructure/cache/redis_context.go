package cache

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"os"
)

var RedisClient *redis.Client
var ctx = context.Background()

var redisKey *string

func InitRedis() {
	var url string
	if os.Getenv("ENV") == "staging" {
		url = os.Getenv("RENDER_REDIS_URL")
		redisKey = &url
		return
	}
	url = os.Getenv("REDIS_URL")
	redisKey = &url
}

func ConnectRedis() *redis.Client {

	opt, err := redis.ParseURL(*redisKey)
	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(opt)

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return RedisClient
}
