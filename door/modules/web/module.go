package web

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/conf"
	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/utils/hs"
)

const moduleName = "web"

type module struct {
	iface.IModule
	*hs.HttpService
}

func New() iface.IModule {
	m := &module{
		IModule:     basic.NewConcurrency(moduleName),
		HttpService: hs.NewHttpService(),
	}
	m.initHandle()
	m.ListenAndServe(conf.Str("web.addr", ""))
	return m
}

func (m *module) Exit() error {
	return m.Stop()
}
