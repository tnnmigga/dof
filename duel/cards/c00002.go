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
}

type cardBasic struct {
	Pairs   []Pair[string, int32]
	Effects []IEffect // 效果
	Skills  []ISkill  // 技能
}

func (c *cardBasic) Value(key string, value ...int32) int32 {
	for _, item := range c.Pairs {
		if item.First != key {
			continue
		}
		if len(value) > 0 {
			item.Second = value[0]
		}
		return item.Second
	}
	if len(value) > 0 {
		c.Pairs = append(c.Pairs, Pair[string, int32]{key, value[0]})
		return value[0]
	}
	return 0
}

func (c *Card) Attack() int32 {
	return c.Value(EnumCardValue.Attack)
}

func (c *Card) Defense() int32 {
	return c.Value(EnumCardValue.Defense)
}

func (c *Card) Health() int32 {
	return c.Value(EnumCardValue.Health)
}



type Card struct {
	cardBasic
}

type IEffect interface {
}

