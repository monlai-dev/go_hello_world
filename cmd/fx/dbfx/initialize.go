package dbfx

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
	"webapp/internal/infrastructure/database"
)

var Module = fx.Provide(
	provideDB,
)

func provideDB() *gorm.DB {
	return database.ConnectDb()
}
