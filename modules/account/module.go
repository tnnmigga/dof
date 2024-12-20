package account

import (
	"eastv2/define"
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
	return define.ModuleName.Account
}

func tokenKey(token string) string {
	return fmt.Sprintf("token:{%s}", token)
}
