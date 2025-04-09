package addressfx

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/services"
)

var Module = fx.Provide(provideAddressService)

func provideAddressService(db *gorm.DB) services.AddressServiceInterface {
	return services.NewAddressService(db)
}
