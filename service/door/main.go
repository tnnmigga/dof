package main

import (
	"eastv2/modules/hss"

	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(
		hss.New(),
	)
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
