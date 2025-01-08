package halebopp

import (
	"github.com/google/wire"
)

type HaleBopp struct {
	Battery
}

var ProvideHaleBopp = wire.Struct(new(HaleBopp), "*")
