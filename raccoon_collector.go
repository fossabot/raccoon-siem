package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/tephrocactus/raccoon-siem/collector"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"log"
	"runtime"
)

type collectorSettings struct {
	ID      string `long:"id" required:"y" description:"collector ID"`
	Core    string `long:"core" default:"http://localhost:7220" description:"core URL"`
	Bus     string `long:"bus" default:"" description:"bus URL"`
	Storage string `long:"storage" default:"" description:"storage URL"`
}

func main() {
	// Parse cmd flags
	settings := new(collectorSettings)
	if _, err := flags.Parse(settings); err != nil {
		log.Fatal(err)
	}

	// Get settings package
	pack := new(sdk.CollectorPackage)
	err := sdk.CoreQuery(settings.Core+"/register/collector/"+settings.ID, pack)
	sdk.PanicOnError(err)

	// Register dictionaries
	err = sdk.RegisterDictionaries(pack.Dictionaries)
	sdk.PanicOnError(err)

	// Register parsers
	registeredParsers, err := sdk.RegisterParsers(pack.Parsers)
	sdk.PanicOnError(err)

	// Registered filters
	registeredFilters, err := sdk.RegisterFilters(pack.Filters)
	sdk.PanicOnError(err)

	// Register destinations
	allDestinationSettings := sdk.GetDefaultDestinationSettings(settings.Storage, settings.Bus)
	allDestinationSettings = append(allDestinationSettings, pack.Destinations...)
	registeredDestinations, err := sdk.RegisterDestinations(allDestinationSettings)
	sdk.PanicOnError(err)

	// Register aggregation rules
	aggregationChannel := make(chan sdk.AggregationChainTask)
	registeredAggregationRules, err := sdk.RegisterAggregationRules(
		pack.AggregationRules,
		pack.AggregationFilters,
		aggregationChannel)
	sdk.PanicOnError(err)

	// Register sources
	parsingChannel := make(chan sdk.ProcessorTask)
	registeredSources, err := sdk.RegisterSources(pack.Sources, parsingChannel)
	sdk.PanicOnError(err)

	// Processor
	proc := collector.Processor{
		ParsingChannel:     parsingChannel,
		AggregationChannel: aggregationChannel,
		Workers:            runtime.NumCPU(),
		Parsers:            registeredParsers,
		Filters:            registeredFilters,
		AggregationRules:   registeredAggregationRules,
		Sources:            registeredSources,
		Destinations:       registeredDestinations,
		ID:                 settings.ID,
		Debug:              pack.Debug,
	}

	proc.Start()
	runtime.Goexit()
}
