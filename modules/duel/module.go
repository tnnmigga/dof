package duel

import (
	"eastv2/define"

	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

type module struct {
	iface.IModule
}

func New() iface.IModule {
	m := &module{
		IModule: basic.NewEventLoop(basic.DefaultMQLen),
	}
	return m
}

func (m module) Name() string {
	return define.ModuleName.Duel
}

func (m *module) Exit() error {
	return nil
}
