package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/tephrocactus/raccoon-siem/correlator"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"log"
	"runtime"
)

type correlatorSettings struct {
	ID            string `long:"id" required:"y" description:"correlator ID"`
	Core          string `long:"core" default:"http://localhost:7220" description:"core URL"`
	Bus           string `long:"bus" default:"nats://localhost:4222" description:"bus URL"`
	Storage       string `long:"storage" default:"http://localhost:9200" description:"storage URL"`
	ALService     string `long:"al.url" default:"localhost:6379" description:"active list service URL"`
	ALServicePool int    `long:"al.pool" default:"10" description:"active list service connection pool size"`
}

func main() {
	// Parse cmd flags
	settings := new(correlatorSettings)
	if _, err := flags.Parse(settings); err != nil {
		log.Fatal(err)
	}

	// Get settings package
	pack := new(sdk.CorrelatorPackage)
	err := sdk.CoreQuery(settings.Core+"/register/correlator/"+settings.ID, pack)
	sdk.PanicOnError(err)

	// Register default event parser
	registeredParsers, err := sdk.RegisterParsers([]sdk.ParserSettings{{Name: "event", Kind: "event", Root: true}})
	sdk.PanicOnError(err)

	// Register active lists and run service client
	alServiceSettings := sdk.ActiveListServiceSettings{
		Name:     sdk.RaccoonActiveListsServiceName,
		PoolSize: settings.ALServicePool,
		URL:      settings.ALService,
	}

	err = sdk.RegisterActiveLists(alServiceSettings, pack.ActiveLists)
	sdk.PanicOnError(err)

	// Register filters
	registeredFilters, err := sdk.RegisterFilters(pack.Filters)
	sdk.PanicOnError(err)

	// Register destinations
	allDestinationSettings := sdk.GetDefaultDestinationSettings(settings.Storage, settings.Bus)
	allDestinationSettings = append(allDestinationSettings, pack.Destinations...)
	registeredDestinations, err := sdk.RegisterDestinations(allDestinationSettings)
	sdk.PanicOnError(err)

	// Register correlation rules
	correlationChainChannel := make(chan sdk.CorrelationChainTask)
	registeredCorrelationRules, err := sdk.RegisterCorrelationRules(
		pack.CorrelationRules,
		registeredFilters,
		correlationChainChannel)
	sdk.PanicOnError(err)

	// Run correlation rules
	sdk.RunCorrelationRules(registeredCorrelationRules)

	// Register default sources
	allSourceSettings := []sdk.SourceSettings{{
		Name:    sdk.RaccoonCorrelationBusName,
		Kind:    sdk.RaccoonCorrelationBusKind,
		Channel: sdk.RaccoonCorrelationBusChannel,
		URL:     settings.Bus,
	}}

	correlationChannel := make(chan sdk.ProcessorTask)
	registeredSources, err := sdk.RegisterSources(allSourceSettings, correlationChannel)
	sdk.PanicOnError(err)

	// Processor
	proc := correlator.Processor{
		CorrelationChannel:      correlationChannel,
		CorrelationChainChannel: correlationChainChannel,
		Workers:                 runtime.NumCPU(),
		Parsers:                 registeredParsers,
		CorrelationRules:        registeredCorrelationRules,
		Sources:                 registeredSources,
		Destinations:            registeredDestinations,
		Debug:                   pack.Debug,
	}

	proc.Start()
	runtime.Goexit()
}
