package hss

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/conf"
	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/utils/hs"
)

type module struct {
	iface.IModule
	*hs.HttpService
}

func New() iface.IModule {
	return &module{
		IModule:     basic.NewConcurrency(),
		HttpService: hs.NewHttpService(),
	}
}

func (m *module) Init() error {
	m.initHandle()
	return m.ListenAndServe(conf.Str("web.addr", ""))
}

func (m module) Name() string {
	return "door"
}

func (m *module) Exit() error {
	return m.Stop()
}
