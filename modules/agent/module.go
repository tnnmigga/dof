package agent

import (
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/conf"
	"github.com/tnnmigga/corev2/iface"
	"github.com/tnnmigga/corev2/log"
)

const (
	ProtocolTCP       = "tcp"
	ProtocolWebSocket = "websocket"
)

type module struct {
	iface.IModule
	protocol string
	listener IListener
	manager  *AgentManager
}

type IListener interface {
	Run() error
	Close()
}

func New() *module {
	protocol := conf.Str("agent.protocol", ProtocolWebSocket)
	m := basic.NewConcurrency()
	a := &module{
		IModule:  m,
		protocol: protocol,
		manager:  NewAgentManager(m),
	}
	switch protocol {
	case ProtocolTCP:
		a.listener = NewTCPListener(a.manager)
	case ProtocolWebSocket:
		a.listener = NewWebSocketListener(a.manager)
	default:
		log.Panicf("unknown agent type: %s", protocol)
	}
	a.register()
	return a
}

func (m module) Name() string {
	return "agent"
}

func (m *module) Run() error {
	return m.listener.Run()
}

func (m *module) Exit() error {
	m.listener.Close()
	for _, agent := range m.manager.agents {
		agent.cancel()
		agent.conn.Close()
	}
	return nil
}
