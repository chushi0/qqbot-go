package module

import (
	"context"
	"log"

	"github.com/chushi0/qqbot-go/src/common/module"
	proto "github.com/chushi0/qqbot-go/src/proto_gen"
)

type API struct {
}

var (
	apiImpl = &API{}
)

func GetAPI() module.API {
	return apiImpl
}

func (API) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	log.Println("SendMessage")
	return &proto.SendMessageResponse{}, nil
}
