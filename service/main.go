package main

import (
	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/basic"
	"github.com/tnnmigga/corev2/conf"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(basic.Create(conf.List[string]("modules")...)...)
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
