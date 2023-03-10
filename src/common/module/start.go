package module

import (
	"github.com/hashicorp/go-plugin"
)

func Start(impl Module) {
	pluginMap := map[string]plugin.Plugin{
		PluginName: &ModuleBridge{Impl: impl},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}
