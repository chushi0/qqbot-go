package main

import (
	"log"

	"github.com/chushi0/qqbot-go/src/common/module"
	proto "github.com/chushi0/qqbot-go/src/proto_gen"
)

type EchoModule struct {
	module.BaseModule
}

func main() {
	module.Start(&EchoModule{})
}

func (*EchoModule) OnLoad(ctx *module.BotContext, req *proto.OnLoadRequest) (resp *proto.EmptyResponse, err error) {
	log.Println("OnLoad")
	rsend, rerr := ctx.SendMessage(ctx, &proto.SendMessageRequest{})
	log.Printf("SendMessage Result: %v, err: %v", rsend, rerr)
	rsend, rerr = ctx.SendMessage(ctx, &proto.SendMessageRequest{})
	log.Printf("SendMessage Result: %v, err: %v", rsend, rerr)
	rsend, rerr = ctx.SendMessage(ctx, &proto.SendMessageRequest{})
	log.Printf("SendMessage Result: %v, err: %v", rsend, rerr)
	resp = &proto.EmptyResponse{}
	return
}
