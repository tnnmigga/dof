package modules

import (
	"eastv2/modules/account"
	"eastv2/modules/agent"
	"eastv2/modules/duel"
	"eastv2/modules/play"
	"eastv2/modules/web"

	"github.com/tnnmigga/corev2/basic"
)

func Init() {
	basic.RegModule(account.New)
	basic.RegModule(agent.New)
	basic.RegModule(play.New)
	basic.RegModule(duel.New)
	basic.RegModule(web.New)
}
