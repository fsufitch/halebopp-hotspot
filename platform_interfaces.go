package halebopp

type ChargeState uint16

const (
	ChargeState_Unknown ChargeState = iota
	ChargeState_Discharging
	ChargeState_Charging
)

type Battery interface {
	Stats() (voltage float64, level float64, err error)
	ChargeState() (ChargeState, error)
}

type Modem interface {
	SignalLevel() (int, error)
}
