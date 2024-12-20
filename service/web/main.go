package main

import (
	"eastv2/modules/door"

	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(
		door.New(),
	)
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
