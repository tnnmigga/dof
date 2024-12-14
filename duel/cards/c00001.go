package cards

import (
	"github.com/tnnmigga/corev2/utils"
)

type DuelContext struct {
	
}

func NewDuelContext() *DuelContext {
	return &DuelContext{}
}

type StateMachineCommand struct {
	Type string
	Args any
}

type StateMachine struct {
	Cmds chan *StateMachineCommand
}

func Next() {
	
}

var EnumCmdType = utils.NewEnum[struct {
	CardDo string
}]()

var EnumEventType = utils.NewEnum[struct {
	BeAttack string
	BeHarm   string
}]()
