package interfaces

import (
	"context"

	"github.com/tommjj/rabbimq-chat/internal/types"
)

type IChatPublish interface {
	Publish(ctx context.Context, mess types.Message) error
}

type IChatConsume interface {
	Run(handler func(mess types.Message) types.AckType) error
}
