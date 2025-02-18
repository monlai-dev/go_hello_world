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
		redisKey = os.Getenv("RENDER_REDIS_URL")
		return
	}

	redisKey = os.Getenv("REDIS_URL")

}

func ConnectRedis() *redis.Client {

	opt, err := redis.ParseURL(redisKey)
	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(opt)

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return RedisClient
}
