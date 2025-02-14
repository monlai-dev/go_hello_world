package rabbitMq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

type RabbitMq struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func ConnectRabbitMq() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		panic(err)
	}

	return conn, nil
}

func NewRabbitClient(conn *amqp.Connection) (*RabbitMq, error) {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return &RabbitMq{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r RabbitMq) DeclareNewQueue(client *RabbitMq, queueName string) error {
	_, err := client.ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r RabbitMq) DeclareNewExchange(client *RabbitMq, exchangeName, exchangeType string) error {
	err := client.ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r RabbitMq) BindQueue(client *RabbitMq, queueName, exchangeName, routingKey string) error {
	err := client.ch.QueueBind(queueName, routingKey, exchangeName, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r RabbitMq) Publish(client *RabbitMq, exchangeName, routingKey string, options amqp.Publishing) error {

	rabbitErr := client.ch.Publish(exchangeName, routingKey, false, false, options)
	if rabbitErr != nil {
		return rabbitErr
	}
	return nil
}

func (r RabbitMq) Consume(queueName string, consumer string) (<-chan amqp.Delivery, error) {
	msgs, err := r.ch.Consume(queueName, consumer, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func Close(client *RabbitMq) {
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing connection")
		}
	}(client.conn)
}
