package modules

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type Battery interface {
	Stats() (voltage float64, level float64, err error)
}

type X278Battery struct {
	bus i2c.BusCloser
}

var _ Battery = &X278Battery{}

const (
	x278_Addr          uint16 = 0x36
	x278_VoltageOffset uint16 = 2
	x278_LevelOffset   uint16 = 4
)

func ProvideX278Battery() (Battery, func(), error) {
	_, err := host.Init()
	if err != nil {
		return nil, func() {}, fmt.Errorf("driver init: %w", err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		return nil, func() {}, fmt.Errorf("bus open: %w", err)
	}

	cleanup := func() {
		bus.Close()
	}

	return &X278Battery{bus: bus}, cleanup, nil
}

func (x278 *X278Battery) Stats() (voltage float64, level float64, err error) {
	i2cData := make([]byte, 8)

	err = x278.bus.Tx(x278_Addr, []byte{0}, i2cData)
	if err != nil {
		return
	}
	fmt.Println(i2cData)

	voltageData, err := reverseEndian(i2cData[1:4])
	if err != nil {
		return
	}
	var voltageRaw uint16
	err = binary.Read(bytes.NewBuffer(voltageData), binary.LittleEndian, &voltageRaw)
	if err != nil {
		return
	}
	voltage = float64(voltageRaw) * 1.25 / 1000 / 16

	levelData, err := reverseEndian(i2cData[4:7])
	if err != nil {
		return
	}

	var levelRaw uint16
	err = binary.Read(bytes.NewBuffer(levelData), binary.BigEndian, &levelRaw)
	if err != nil {
		return
	}
	level = float64(levelRaw) / 256

	return
}

func reverseEndian(data []byte) ([]byte, error) {
	buf1 := bytes.Buffer{}
	err := binary.Write(&buf1, binary.BigEndian, data)
	if err != nil {
		return nil, err
	}

	buf2 := bytes.Buffer{}
	err = binary.Write(&buf2, binary.LittleEndian, buf1.Bytes())
	if err != nil {
		return nil, err
	}

	return buf2.Bytes(), nil
}
