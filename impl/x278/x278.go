package x278

import (
	"github.com/fsufitch/halebopp-hotspot"
	"github.com/google/wire"
)

type X278 struct {
	I2CBus
	GPIO
}

var ProvideX278 = wire.NewSet(
	NewI2CBus,
	NewGPIO,
	wire.Struct(new(X278), "*"),
	wire.Bind(new(halebopp.Battery), new(X278)),
)
