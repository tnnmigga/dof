package duel

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

const moduleName = "duel"

type module struct {
	iface.IModule
}

func New() iface.IModule {
	m := &module{
		IModule: basic.NewEventLoop(moduleName, basic.DefaultMQLen),
	}
	return m
}

func (m *module) Exit() error {
	return nil
}
