package dummy

type Modem struct {
	Signal `yaml:"signal"`
}

type Signal struct {
	Level int    `yaml:"level"`
	Error string `yaml:"error"`
}

var defaultModem = Modem{
	Signal: Signal{
		Level: 4,
		Error: "",
	},
}

func (sl Signal) SignalLevel() (int, error) {
	return sl.Level, newDummyError(sl.Error)
}
