package web

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

func New() *module {
	m := &module{
		IModule:     basic.NewConcurrency(),
		HttpService: hs.NewHttpService(),
	}
	m.initHandle()
	m.ListenAndServe(conf.Str("web.addr", ""))
	return m
}

func (m module) Name() string {
	return "web"
}

func (m *module) Exit() error {
	return m.Stop()
}
