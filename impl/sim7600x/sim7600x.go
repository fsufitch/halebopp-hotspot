package sim7600x

import (
	"github.com/fsufitch/halebopp-hotspot"
	"github.com/google/wire"
)

type SIM7600X struct{}

func (s SIM7600X) SignalLevel() (int, error) {
	return -1, nil
}

var ProvideModem = wire.NewSet(
	wire.Struct(new(SIM7600X), "*"),
	wire.Bind(new(halebopp.Modem), new(SIM7600X)),
)
