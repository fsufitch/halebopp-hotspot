package x278

import (
	"github.com/fsufitch/halebopp-hotspot"
	"github.com/google/wire"
)

type X278 struct {
	I2CBus
	GPIO
}

func ProvideBattery(x278 X278) halebopp.Battery {
	return x278
}

var ProvideX278 = wire.NewSet(
	wire.Struct(new(X278), "*"),
	ProvideBattery,
	NewI2CBus,
	NewGPIO,
)
