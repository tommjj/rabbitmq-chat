package publish

import (
	"context"
	"testing"
	"time"

	"github.com/tommjj/rabbimq-chat/internal/rabbitmq"
	"github.com/tommjj/rabbimq-chat/internal/types"
)

const connectionString = "amqp://guest:guest@localhost:5672/"

func TestPub(t *testing.T) {
	conn, err := rabbitmq.NewConn(connectionString, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	chatPub, err := NewChatPub(conn)
	if err != nil {
		t.Fatal(err)
	}

	chatPub.Publish(context.Background(), types.Message{
		From: types.User{Name: "laplili"},
		Text: "Hello world!",
	})

	time.Sleep(time.Second)
}
