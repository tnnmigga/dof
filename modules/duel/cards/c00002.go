package cards

import (
	"fmt"
)

var factory = map[int32]func() ICard{}

func RegCard(gen func() ICard) {
	card := gen()
	cid := card.ID()
	if _, ok := factory[cid]; ok {
		panic(fmt.Errorf("card %d is exists", cid))
	}
	factory[cid] = gen
}

func GenCard(cid int32) ICard {
	return factory[cid]()
}

type ICard interface {
	ID() int32
	Attack() int32
	Defense() int32
	Health() int32
}

type Card struct {
	attack  int32
	defense int32
	health  int32
	Effects []IEffect // 效果
	Skills  []ISkill  // 技能
}

func (c *Card) Attack() int32 {
	return c.attack
}

func (c *Card) Defense() int32 {
	return c.defense
}

func (c *Card) Health() int32 {
	return c.health
}

type IEffect interface {
}

type ISkill interface{}
