package module

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/chushi0/qqbot-go/src/common/module"
	proto "github.com/chushi0/qqbot-go/src/proto_gen"

	"github.com/hashicorp/go-plugin"
)

type ModuleManager struct {
	Modules []*Module
}

func (*ModuleManager) LoadModule(name string) error {
	pluginMap := map[string]plugin.Plugin{
		module.PluginName: &module.ModuleBridge{},
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  module.HandshakeConfig,
		Plugins:          pluginMap,
		Cmd:              exec.Command(ModuleDirectory + name + "/bin/main"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})

	rpcClient, err := client.Client()
	if err != nil {
		return err
	}
	raw, err := rpcClient.Dispense(module.PluginName)
	if err != nil {
		return err
	}

	module := raw.(*module.ModuleRPC)
	module.StartServer(GetAPI())
	for i := 0; i < 3; i++ {
		resp, err := module.OnLoad(context.TODO(), &proto.OnLoadRequest{})
		fmt.Println(resp)
		fmt.Println(err)
	}

	module.CloseServer()
	client.Kill()

	return nil
}
