package cards

type C10001 struct {
	Card
}

func NewC10001() ICard {
	card := &C10001{}
	card.attack = 1000
	card.defense = 1000
	card.health = 1000
	return card
}

func (c C10001) ID() int32 {
	return 10001
}
