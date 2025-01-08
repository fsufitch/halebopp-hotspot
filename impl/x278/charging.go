package x278

import (
	"github.com/fsufitch/halebopp-hotspot"
	"github.com/warthog618/go-gpiocdev"
)

const x278_GPIO_PowerConnected = "GPIO6"

type GPIO struct {
	gpio6 *gpiocdev.Line
}

func NewGPIO() (gpio GPIO, err error) {
	chip, offset, err := gpiocdev.FindLine("GPIO6")
	if err != nil {
		return
	}

	gpio.gpio6, err = gpiocdev.RequestLine(chip, offset, gpiocdev.AsInput)
	if err != nil {
		return
	}

	return
}

func (x GPIO) ChargeState() (halebopp.ChargeState, error) {
	value, err := x.gpio6.Value()
	if err != nil {
		return halebopp.ChargeState_Unknown, err
	}
	if value > 0 {
		return halebopp.ChargeState_Discharging, nil
	}
	return halebopp.ChargeState_Charging, nil

}
