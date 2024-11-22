package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tommjj/rabbitmq-chat/internal/chat"
	"github.com/tommjj/rabbitmq-chat/internal/rabbitmq"
	"github.com/tommjj/rabbitmq-chat/internal/rabbitmq/consume"
	"github.com/tommjj/rabbitmq-chat/internal/rabbitmq/publish"
	"github.com/tommjj/rabbitmq-chat/internal/x/types"
	"github.com/tommjj/rabbitmq-chat/pkg/pcolor"
)

const connectionString = "amqp://guest:guest@localhost:5672/"

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := rabbitmq.NewConn(connectionString, nil)
	fatalOnErr(err)
	defer conn.Close()

	fmt.Println("enter your name?")
	name := chat.GetInput()
	if len(name) == 0 {
		log.Fatal("name is invalid")
	}

	publisher, err := publish.NewChatPub(conn)
	fatalOnErr(err)

	chatConsume, err := consume.NewChatConsume(conn, name)
	fatalOnErr(err)

	chatConsume.Run(func(mess types.Message) types.AckType {
		defer fmt.Print("> ")
		if mess.From.Name == name {
			pcolor.Yellow.Printf("%s: %s\n", mess.From.Name, mess.Text)
		} else {
			pcolor.Blue.Printf("%s: %s\n", mess.From.Name, mess.Text)
		}

		return types.Ack
	})

	pcolor.Green.Println("start chat")
	for {
		mess := chat.GetInput()
		if mess == "exit()" {
			break
		}

		err = publisher.Publish(context.Background(), types.Message{
			From: types.User{Name: name},
			Text: mess,
		})
		fatalOnErr(err)
	}
}
