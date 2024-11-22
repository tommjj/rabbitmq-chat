package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConn struct {
	*amqp.Connection
}

func NewConn(connStr string, config *amqp.Config) (*RabbitMQConn, error) {
	var conn *amqp.Connection
	var err error

	if config == nil {
		conn, err = amqp.Dial(connStr)
	} else {
		conn, err = amqp.DialConfig(connStr, *config)
	}

	if err != nil {
		return nil, err
	}

	return &RabbitMQConn{conn}, nil
}
