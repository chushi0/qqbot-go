package module

import (
	"context"

	proto "github.com/chushi0/qqbot-go/src/proto_gen"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type Module interface {
	OnLoad(*BotContext, *proto.OnLoadRequest) (*proto.EmptyResponse, error)
	OnUnload(*BotContext, *proto.OnUnloadRequest) (*proto.EmptyResponse, error)
	OnMessage(*BotContext, *proto.OnMessageRequest) (*proto.EmptyResponse, error)
	OnImcCall(*BotContext, *proto.OnImcCallRequest) (*proto.EmptyResponse, error)
}

type BaseModule struct {
}

func (*BaseModule) OnLoad(*BotContext, *proto.OnLoadRequest) (*proto.EmptyResponse, error) {
	return &proto.EmptyResponse{}, nil
}

func (*BaseModule) OnUnload(*BotContext, *proto.OnUnloadRequest) (*proto.EmptyResponse, error) {
	return &proto.EmptyResponse{}, nil
}

func (*BaseModule) OnMessage(*BotContext, *proto.OnMessageRequest) (*proto.EmptyResponse, error) {
	return &proto.EmptyResponse{}, nil
}

func (*BaseModule) OnImcCall(*BotContext, *proto.OnImcCallRequest) (*proto.EmptyResponse, error) {
	return &proto.EmptyResponse{}, nil
}

type ModuleBridge struct {
	plugin.Plugin
	Impl Module
}

func (p *ModuleBridge) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterModuleServer(s, &moduleServer{impl: p.Impl, broker: broker})
	return nil
}

func (ModuleBridge) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &ModuleRPC{client: proto.NewModuleClient(c), broker: broker}, nil
}
