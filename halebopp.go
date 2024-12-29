package halebopp

import (
	"github.com/fsufitch/halebopp-hotspot/modules"
	"github.com/google/wire"
)

type HaleBopp struct {
	Battery  modules.Battery
	Charging modules.Charging
}

var ProvideHaleBopp = wire.Struct(new(HaleBopp), "*")
