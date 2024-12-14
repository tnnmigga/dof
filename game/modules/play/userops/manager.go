package userops

import (
	"eastv2/game/modules/play/userops/userdata"

	"github.com/tnnmigga/corev2/iface"
)

var manager *Manager

type Manager struct {
	iface.IModule
	cache   map[uint64]*userdata.Meta
	waiting map[uint64][]func(*userdata.Meta, error)
}

func Init(m iface.IModule) {
	manager = &Manager{
		IModule: m,
		cache:   map[uint64]*userdata.Meta{},
		waiting: map[uint64][]func(*userdata.Meta, error){},
	}
	manager.regHandle()
}
