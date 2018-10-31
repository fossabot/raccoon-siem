package logger

import (
	"github.com/tephrocactus/raccoon-siem/sdk"
)

type Factory struct {
	destinations []sdk.IDestination
}

func (r *Factory) NewInstance(name string, level ...int) *Instance {
	return NewInstance(name, r.destinations, level...)
}

func (r *Factory) initDestination(dstSettings sdk.DestinationSettings) error {
	dst, err := sdk.NewDestination(dstSettings)
	if err != nil {
		return err
	}

	if err = dst.Run(); err != nil {
		return err
	}

	r.destinations = append(r.destinations, dst)
	return nil
}

func NewFactory(settings []sdk.DestinationSettings) (*Factory, error) {
	factory := new(Factory)
	for i := range settings {
		if err := factory.initDestination(settings[i]); err != nil {
			return nil, err
		}
	}
	return factory, nil
}
