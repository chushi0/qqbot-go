package module

import (
	"context"

	proto "github.com/chushi0/qqbot-go/src/proto_gen"

	"github.com/hashicorp/go-plugin"
)

type moduleServer struct {
	proto.UnimplementedModuleServer
	impl   Module
	broker *plugin.GRPCBroker
}

func (m *moduleServer) OnLoad(ctx context.Context, req *proto.OnLoadRequest) (*proto.EmptyResponse, error) {
	conn, err := m.broker.Dial(req.Base.GetServer())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	botCtx := &BotContext{
		Context:   ctx,
		APIClient: proto.NewAPIClient(conn),
	}
	return m.impl.OnLoad(botCtx, req)
}

func (m *moduleServer) OnUnload(ctx context.Context, req *proto.OnUnloadRequest) (*proto.EmptyResponse, error) {
	conn, err := m.broker.Dial(req.Base.GetServer())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	botCtx := &BotContext{
		Context:   ctx,
		APIClient: proto.NewAPIClient(conn),
	}
	return m.impl.OnUnload(botCtx, req)
}

func (m *moduleServer) OnMessage(ctx context.Context, req *proto.OnMessageRequest) (*proto.EmptyResponse, error) {
	conn, err := m.broker.Dial(req.Base.GetServer())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	botCtx := &BotContext{
		Context:   ctx,
		APIClient: proto.NewAPIClient(conn),
	}
	return m.impl.OnMessage(botCtx, req)
}

func (m *moduleServer) OnImcCall(ctx context.Context, req *proto.OnImcCallRequest) (*proto.EmptyResponse, error) {
	conn, err := m.broker.Dial(req.Base.GetServer())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	botCtx := &BotContext{
		Context:   ctx,
		APIClient: proto.NewAPIClient(conn),
	}
	return m.impl.OnImcCall(botCtx, req)
}
