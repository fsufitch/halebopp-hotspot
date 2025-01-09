package dummy

import "github.com/fsufitch/halebopp-hotspot"

type Battery struct {
	BatteryLevel    `yaml:"level"`
	BatteryCharging `yaml:"charging"`
}

type BatteryLevel struct {
	Voltage float64 `yaml:"voltage"`
	Level   float64 `yaml:"level"`
	Error   string  `yaml:"error"`
}

type BatteryCharging struct {
	Charging halebopp.ChargeState `yaml:"charging"`
	Error    string               `yaml:"error"`
}

var defaultBattery = Battery{
	BatteryLevel: BatteryLevel{
		Voltage: 9.876,
		Level:   0.9,
		Error:   "",
	},
	BatteryCharging: BatteryCharging{
		Charging: halebopp.ChargeState_Charging,
		Error:    "",
	},
}

func (bl BatteryLevel) Stats() (voltage float64, level float64, err error) {
	return bl.Voltage, bl.Level, newDummyError(bl.Error)
}

func (bc BatteryCharging) ChargeState() (halebopp.ChargeState, error) {
	return bc.Charging, newDummyError(bc.Error)
}
