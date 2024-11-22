package consume

import (
	"testing"
	"time"

	"github.com/tommjj/rabbimq-chat/internal/interfaces"
	"github.com/tommjj/rabbimq-chat/internal/rabbitmq"
	"github.com/tommjj/rabbimq-chat/internal/types"
	"github.com/tommjj/rabbimq-chat/pkg/pcolor"
)

const connectionString = "amqp://guest:guest@localhost:5672/"

func TestChatConsume(t *testing.T) {
	conn, err := rabbitmq.NewConn(connectionString, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	consume, err := NewChatConsume(conn, "laplala")
	if err != nil {
		t.Fatal(err)
	}

	consume.Run(func(mess types.Message) interfaces.AckType {
		t.Log(pcolor.Blue.Sprint(mess))
		return interfaces.NackRequeue
	})

	time.Sleep(time.Hour)
}
