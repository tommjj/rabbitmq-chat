package publish

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tommjj/rabbimq-chat/internal/rabbitmq"
	"github.com/tommjj/rabbimq-chat/internal/types"
)

type ChatPub struct {
	conn *rabbitmq.RabbitMQConn
	ch   *amqp.Channel
}

func NewChatPub(conn *rabbitmq.RabbitMQConn) (*ChatPub, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &ChatPub{
		conn: conn,
		ch:   ch,
	}, nil
}

func (c *ChatPub) Publish(ctx context.Context, mess types.Message) error {
	bytes, err := json.Marshal(mess)
	if err != nil {
		return err
	}

	err = c.ch.PublishWithContext(ctx, rabbitmq.ChatTopicExchange, "chat."+mess.From.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        bytes,
	})
	return err
}
