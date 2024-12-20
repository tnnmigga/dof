package account

import (
	"fmt"

	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

type module struct {
	iface.IModule
}

func New() iface.IModule {
	m := &module{
		IModule: basic.NewConcurrency(),
	}
	return m
}

func (m module) Name() string {
	return "account"
}

func tokenKey(token string) string {
	return fmt.Sprintf("token:{%s}", token)
}
