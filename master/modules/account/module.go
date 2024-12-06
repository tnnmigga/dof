package account

import (
	"fmt"

	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/iface"
)

const moduleName = "account"

type module struct {
	iface.IModule
}

func New() iface.IModule {
	m := &module{
		IModule: basic.NewConcurrency(moduleName),
	}
	m.Init()
	return m
}

func tokenKey(token string) string {
	return fmt.Sprintf("token:{%s}", token)
}
