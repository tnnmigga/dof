package main

import (
	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
