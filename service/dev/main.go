package main

import (
	"eastv2/modules/account"
	"eastv2/modules/agent"
	"eastv2/modules/duel"
	"eastv2/modules/hss"
	"eastv2/modules/play"

	"github.com/tnnmigga/corev2"
	"github.com/tnnmigga/corev2/system"
)

func main() {
	app := corev2.DefaultApp()
	app.Append(
		account.New(),
		agent.New(),
		play.New(),
		duel.New(),
		hss.New(),
	)
	app.Launch()
	defer app.Shutdown()
	system.WaitExitSignal()
}
