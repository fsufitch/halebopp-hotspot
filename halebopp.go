package halebopp

import (
	"github.com/google/wire"
)

type HaleBopp struct {
	Battery
	Modem
}

var ProvideHaleBopp = wire.Struct(new(HaleBopp), "*")
