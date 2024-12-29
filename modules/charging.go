package modules

import (
	"github.com/warthog618/go-gpiocdev"
)

type ChargeState uint16

const (
	ChargeState_Unknown ChargeState = iota
	ChargeState_Discharging
	ChargeState_Charging
)

type Charging interface {
	State() (ChargeState, error)
}

const x278_GPIO_PowerConnected = "GPIO6"

type X278Charging struct {
	line *gpiocdev.Line
}

func ProvideX278Charging() (Charging, error) {
	chip, offset, err := gpiocdev.FindLine("GPIO6")
	if err != nil {
		return nil, err
	}

	line, err := gpiocdev.RequestLine(chip, offset, gpiocdev.AsInput)
	if err != nil {
		return nil, err
	}

	return &X278Charging{line: line}, nil
}

func (x *X278Charging) State() (ChargeState, error) {

	value, err := x.line.Value()
	if err != nil {
		return ChargeState_Unknown, err
	}
	if value > 0 {
		return ChargeState_Discharging, nil
	}
	return ChargeState_Charging, nil

}

var _ Charging = &X278Charging{}
