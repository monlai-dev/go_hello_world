package redisfx

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"webapp/internal/infrastructure/cache"
)

var Module = fx.Provide(provideRedisClient)

func provideRedisClient() *redis.Client {
	return cache.ConnectRedis()
}
