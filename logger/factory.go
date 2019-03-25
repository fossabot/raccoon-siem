package logger

import (
	"github.com/tephrocactus/raccoon-siem/sdk/destinations"
)

type Factory struct {
	destinations []destinations.IDestination
}

func (r *Factory) NewInstance(name string, level ...int) *Instance {
	return NewInstance(name, r.destinations, level...)
}

func (r *Factory) initDestination(cfg destinations.Config) error {
	dst, err := destinations.New(cfg)
	if err != nil {
		return err
	}

	if err = dst.Start(); err != nil {
		return err
	}

	r.destinations = append(r.destinations, dst)
	return nil
}

func NewFactory(cfg ...destinations.Config) (*Factory, error) {
	if len(cfg) == 0 {
		cfg = []destinations.Config{{Kind: destinations.DestinationConsole}}
	}

	factory := new(Factory)
	for i := range cfg {
		if err := factory.initDestination(cfg[i]); err != nil {
			return nil, err
		}
	}

	return factory, nil
}
