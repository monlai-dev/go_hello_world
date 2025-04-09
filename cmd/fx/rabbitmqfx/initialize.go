package rabbitmqfx

import (
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"webapp/internal/infrastructure/rabbitMq"
)

var Module = fx.Provide(
	provideRabbitConnection,
	provideRabbitClient)

func provideRabbitConnection() (*amqp091.Connection, error) {
	return rabbitMq.ConnectRabbitMq()
}

func provideRabbitClient(conn *amqp091.Connection) (*rabbitMq.RabbitMq, error) {
	return rabbitMq.NewRabbitClient(conn)
}
