package module

import (
	"context"

	proto "github.com/chushi0/qqbot-go/src/proto_gen"
)

type BotContext struct {
	context.Context
	proto.APIClient
}
