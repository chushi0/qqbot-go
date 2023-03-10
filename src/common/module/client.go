package module

import (
	"context"

	proto "github.com/chushi0/qqbot-go/src/proto_gen"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type ModuleRPC struct {
	proto.UnimplementedModuleServer
	client   proto.ModuleClient
	broker   *plugin.GRPCBroker
	serverID uint32
	close    chan int
	start    chan int
}

func (m *ModuleRPC) StartServer(api API) {
	apiServer := &GRPCAPIServer{Impl: api}
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s := grpc.NewServer(opts...)
		proto.RegisterAPIServer(s, apiServer)

		return s
	}

	brokerID := m.broker.NextId()
	m.serverID = brokerID
	m.close = make(chan int, 1)
	m.start = make(chan int, 1)
	go func() {
		for {
			select {
			case <-m.close:
				return
			case <-m.start:
				go m.broker.AcceptAndServe(brokerID, serverFunc)
			}
		}
	}()
}

func (m *ModuleRPC) CloseServer() {
	m.close <- 1
	m.broker.Close()
}

func (m *ModuleRPC) OnLoad(ctx context.Context, req *proto.OnLoadRequest) (*proto.EmptyResponse, error) {
	req.Base = m.createBase(req.Base)
	m.start <- 1
	return m.client.OnLoad(ctx, req)
}

func (m *ModuleRPC) OnUnload(ctx context.Context, req *proto.OnUnloadRequest) (*proto.EmptyResponse, error) {
	req.Base = m.createBase(req.Base)
	m.start <- 1
	return m.client.OnUnload(ctx, req)
}

func (m *ModuleRPC) OnMessage(ctx context.Context, api API, req *proto.OnMessageRequest) (*proto.EmptyResponse, error) {
	req.Base = m.createBase(req.Base)
	m.start <- 1
	return m.client.OnMessage(ctx, req)
}

func (m *ModuleRPC) OnImcCall(ctx context.Context, api API, req *proto.OnImcCallRequest) (*proto.EmptyResponse, error) {
	req.Base = m.createBase(req.Base)
	m.start <- 1
	return m.client.OnImcCall(ctx, req)
}

func (m *ModuleRPC) createBase(base *proto.Base) *proto.Base {
	if base == nil {
		base = &proto.Base{}
	}
	base.Server = m.serverID
	return base
}
