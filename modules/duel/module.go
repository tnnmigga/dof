package duel

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

type module struct {
	iface.IModule
}

func New() *module {
	m := &module{
		IModule: basic.NewEventLoop(basic.DefaultMQLen),
	}
	return m
}

func (m module) Name() string {
	return "duel"
}

func (m *module) Exit() error {
	return nil
}
