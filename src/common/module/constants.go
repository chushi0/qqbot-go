package module

import (
	"errors"

	"github.com/hashicorp/go-plugin"
)

const (
	PluginName      string = "default"
	ProtocolVersion uint   = 1
)

var (
	HandshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  ProtocolVersion,
		MagicCookieKey:   "K",
		MagicCookieValue: "V",
	}

	ErrImcNotFound = errors.New("target imc func not found")
)
