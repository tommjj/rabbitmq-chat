package consume

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tommjj/rabbitmq-chat/internal/rabbitmq"
	"github.com/tommjj/rabbitmq-chat/internal/x/interfaces"
	"github.com/tommjj/rabbitmq-chat/internal/x/types"
)

type ChatConsume struct {
	queueName string
	conn      *rabbitmq.RabbitMQConn
	ch        *amqp.Channel
}

func NewChatConsume(conn *rabbitmq.RabbitMQConn, username string) (interfaces.IChatConsume, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queueName := rabbitmq.ChatPrefix + "." + username
	_, err = ch.QueueDeclare(
		queueName,
		false,
		true,
		true,
		false,
		amqp.Table{
			"x-dead-letter-exchange": rabbitmq.DeadChatExchange,
		},
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(queueName, rabbitmq.ChatPrefix+".*", rabbitmq.ChatTopicExchange, false, nil)
	if err != nil {
		return nil, err
	}

	return &ChatConsume{
		queueName: queueName,
		conn:      conn,
		ch:        ch,
	}, nil
}

func (c *ChatConsume) Run(handler func(mess types.Message) types.AckType) error {
	deliveryCh, err := c.ch.Consume(c.queueName, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range deliveryCh {
			var body types.Message
			err = json.Unmarshal(msg.Body, &body)
			if err != nil {
				msg.Nack(false, false)
				continue
			}

			ackType := handler(body)
			switch ackType {
			case types.Ack:
				msg.Ack(false)
			case types.NackRequeue:
				msg.Nack(false, true)
			case types.NackDiscard:
				msg.Nack(false, false)
			}
		}
	}()

	return nil
}
