package define

import "github.com/tnnmigga/corev2/utils"

var ModuleName = utils.NewEnum[struct {
	Play    string
	Account string
	Agent   string
	Duel    string
	Door    string
}]()
