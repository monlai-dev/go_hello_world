package websocketfx

import (
	"go.uber.org/fx"
	"webapp/internal/services"
)

var Module = fx.Provide(provideWebsocketService)

func provideWebsocketService() *services.WebsocketService {
	return services.NewWebsocketService()
}
