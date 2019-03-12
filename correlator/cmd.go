package correlator

import (
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"runtime"
)

var (
	Cmd = &cobra.Command{
		Use:   "correlator",
		Short: "start raccoon correlator",
		Args:  cobra.ExactArgs(0),
		RunE:  run,
	}

	// String flags variables
	coreURL, busURL, storageURL, activeListSvcURL, correlatorID, metricsPort string

	// Int flags variables
	activeListSvcPoolSize int

	// Bool flags variables
	debugMode bool
)

func init() {
	// Raccoon core URL
	Cmd.Flags().StringVarP(
		&coreURL,
		"core",
		"c",
		"http://localhost:7220",
		"raccoon core URL")

	// Raccoon bus URL
	Cmd.Flags().StringVarP(
		&busURL,
		"bus",
		"b",
		"nats://localhost:4222",
		"raccoon bus URL")

	// Raccoon storage URL
	Cmd.Flags().StringVarP(
		&storageURL,
		"storage",
		"s",
		"http://localhost:9200",
		"raccoon storage URL")

	// Raccoon correlator ID
	Cmd.Flags().StringVarP(
		&correlatorID,
		"id",
		"i",
		"",
		"raccoon correlator ID")

	// Raccoon active list service URL
	Cmd.Flags().StringVarP(
		&activeListSvcURL,
		"al.url",
		"a",
		"localhost:6379",
		"raccoon active list service URL")

	// Raccoon active list service connection pool size
	Cmd.Flags().StringVarP(
		&activeListSvcURL,
		"al.pool",
		"",
		"32",
		"raccoon active list service connection pool size")

	// Prometheus metrics port
	Cmd.Flags().StringVarP(
		&metricsPort,
		"metrics",
		"m",
		"7221",
		"prometheus metrics port")

	// Debug mode
	Cmd.Flags().BoolVarP(
		&debugMode,
		"debug",
		"d",
		false,
		"debug mode")

	Cmd.MarkFlagRequired("id")
}

func run(_ *cobra.Command, _ []string) error {
	// Get settings package
	pack := new(sdk.CorrelatorPackage)
	err := sdk.CoreQuery(coreURL+"/register/correlator/"+correlatorID, pack)
	sdk.PanicOnError(err)

	// Register default event parser
	registeredParsers, err := sdk.RegisterParsers([]sdk.ParserSettings{{Name: "event", Kind: "event", Root: true}})
	sdk.PanicOnError(err)

	// Register active lists and run service client
	alServiceSettings := sdk.ActiveListServiceSettings{
		Name:     sdk.RaccoonActiveListsServiceName,
		PoolSize: activeListSvcPoolSize,
		URL:      activeListSvcURL,
	}

	err = sdk.RegisterActiveLists(alServiceSettings, pack.ActiveLists)
	sdk.PanicOnError(err)

	// Register filters
	registeredFilters, err := sdk.RegisterFilters(pack.Filters)
	sdk.PanicOnError(err)

	// Register destinations
	allDestinationSettings := sdk.GetDefaultDestinationSettings(storageURL, busURL, debugMode)
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

	// Register default connectors
	allConnectorConfigs := []sdk.Config{{
		Name:    sdk.RaccoonCorrelationBusName,
		Kind:    sdk.RaccoonCorrelationBusKind,
		Subject: sdk.RaccoonCorrelationBusChannel,
		URL:     busURL,
	}}

	correlationChannel := make(chan *sdk.ProcessorTask)
	registeredConnectors, err := sdk.RegisterConnectors(allConnectorConfigs, correlationChannel)
	sdk.PanicOnError(err)

	// Processor
	proc := Processor{
		CorrelationChannel:      correlationChannel,
		CorrelationChainChannel: correlationChainChannel,
		Workers:                 runtime.NumCPU(),
		Parsers:                 registeredParsers,
		CorrelationRules:        registeredCorrelationRules,
		Connectors:              registeredConnectors,
		Destinations:            registeredDestinations,
		Debug:                   debugMode,
	}

	sdk.PrintConfiguration(
		registeredConnectors,
		registeredParsers,
		registeredCorrelationRules,
		pack.ActiveLists,
		registeredDestinations)

	if err := proc.Start(); err != nil {
		return err
	}

	runtime.Goexit()
	return nil
}
