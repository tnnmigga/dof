package cards

import "github.com/tnnmigga/corev2/utils"

var EnumCardValue = utils.NewEnum[struct {
	Attack  string
	Defense string
	Health  string
}]()

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}
