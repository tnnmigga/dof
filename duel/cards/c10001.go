package cards

type C10001 struct {
	Card
}

func NewC10001() ICard {
	card := &C10001{}
	card.Value(EnumCardValue.Attack, 1000)
	card.Value(EnumCardValue.Defense, 1000)
	card.Value(EnumCardValue.Health, 1000)
	return card
}

func (c C10001) ID() int32 {
	return 10001
}
