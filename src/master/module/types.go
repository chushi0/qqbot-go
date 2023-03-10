package module

import (
	"os/exec"
)

type IModule interface {
}

type Module struct {
	*exec.Cmd
	ModuleConf
}

type ModuleConf struct {
	ID          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Priority    uint32 `yaml:"priority"`
}
