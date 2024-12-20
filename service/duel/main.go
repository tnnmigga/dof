package main

import (
	"eastv2/modules/duel"

	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(
		duel.New(),
	)
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
