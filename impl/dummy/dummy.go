package dummy

import (
	"io"
	"log"
	"os"

	"github.com/fsufitch/halebopp-hotspot"
	"github.com/google/wire"
	"gopkg.in/yaml.v3"
)

type Dummy struct {
	Battery `yaml:"battery"`
	Modem   `yaml:"modem"`
}

var defaultDummy = Dummy{
	Battery: defaultBattery,
	Modem:   defaultModem,
}

func ReadDummy() (dummy Dummy, err error) {
	dummyYAMLPath := os.Getenv("HALEBOPP_DUMMY_YAML")
	if dummyYAMLPath == "" {
		log.Print("dummy data file unset (HALEBOPP_DUMMY_YAML), using defaults")
		return defaultDummy, nil
	}
	log.Printf("reading dummy data from: %s", dummyYAMLPath)

	fp, err := os.Open(dummyYAMLPath)
	defer fp.Close()
	if err != nil {
		return
	}
	data, err := io.ReadAll(fp)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &dummy)
	return
}

var ProvideDummyImplementations = wire.NewSet(
	ReadDummy,
	wire.Bind(new(halebopp.Battery), new(Dummy)),
	wire.Bind(new(halebopp.Modem), new(Dummy)),
)
