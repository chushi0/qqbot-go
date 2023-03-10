package module

import (
	"context"

	proto "github.com/chushi0/qqbot-go/src/proto_gen"
)

type API interface {
	SendMessage(context.Context, *proto.SendMessageRequest) (*proto.SendMessageResponse, error)
}

type GRPCAPIServer struct {
	proto.UnimplementedAPIServer
	Impl API
}

func (a *GRPCAPIServer) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.SendMessageResponse, error) {
	return a.Impl.SendMessage(ctx, req)
}
