package play

import (
	"eastv2/modules/play/domain"
	"eastv2/modules/play/domain/impl"
	"eastv2/modules/play/userops"

	"github.com/tnnmigga/corev2/basic"
)

type module struct {
	*domain.Domain
}

func New() *module {
	m := &module{
		Domain: domain.New(basic.NewEventLoop(basic.DefaultMQLen)),
	}
	userops.Init(m)
	impl.Init(m.Domain)
	return m
}

func (m module) Name() string {
	return "play"
}
