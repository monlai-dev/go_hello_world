package accountfx

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/repositories"
	"webapp/internal/services"
)

var Module = fx.Provide(
	provideAccountRepository,
	provideAccountService)

func provideAccountRepository(db *gorm.DB) repositories.AccountRepositoryInterface {
	return repositories.NewAccountRepository(db)
}

func provideAccountService(db *gorm.DB, accountRepository repositories.AccountRepositoryInterface, redisClient *redis.Client, addressService services.AddressServiceInterface) services.AccountServiceInterface {
	return services.NewAccountService(db, addressService, redisClient, accountRepository)
}
