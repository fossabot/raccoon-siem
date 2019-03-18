package collector

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "collector",
		Short: "start raccoon collector",
		Args:  cobra.ExactArgs(0),
		RunE:  run,
	}

	// String flags variables
	coreURL, busURL, storageURL, collectorID, metricsPort string

	// Bool flags variables
	printRaw, debugMode bool
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
		"",
		"raccoon bus URL")

	// Raccoon storage URL
	Cmd.Flags().StringVarP(
		&storageURL,
		"storage",
		"s",
		"",
		"raccoon storage URL")

	// Raccoon collector ID
	Cmd.Flags().StringVarP(
		&collectorID,
		"id",
		"i",
		"",
		"raccoon collector ID")

	// Prometheus metrics port
	Cmd.Flags().StringVarP(
		&metricsPort,
		"metrics",
		"m",
		"7221",
		"prometheus metrics port")

	// Print raw messages
	Cmd.Flags().BoolVarP(
		&printRaw,
		"raw",
		"r",
		false,
		"print raw input messages")

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
	//// Get settings package
	//pack := new(sdk.CollectorPackage)
	//if err := sdk.CoreQuery(coreURL+"/register/collector/"+collectorID, pack); err != nil {
	//	return err
	//}
	//
	//// Register dictionaries
	//if err := sdk.RegisterDictionaries(pack.Dictionaries); err != nil {
	//	return err
	//}
	//
	//// Register parsers
	//registeredParsers, err := sdk.RegisterParsers(pack.Parsers)
	//if err != nil {
	//	return err
	//}
	//
	//// Registered filters
	//registeredFilters, err := sdk.RegisterFilters(pack.Filters)
	//if err != nil {
	//	return err
	//}
	//
	//// Register destinations
	//allDestinationSettings := sdk.GetDefaultDestinationSettings(storageURL, busURL, debugMode)
	//allDestinationSettings = append(allDestinationSettings, pack.Destinations...)
	//registeredDestinations, err := sdk.RegisterDestinations(allDestinationSettings)
	//if err != nil {
	//	return err
	//}
	//
	//// Register aggregation rules
	//aggregationChannel := make(chan sdk.AggregationChainTask)
	//registeredAggregationRules, err := sdk.RegisterAggregationRules(
	//	pack.AggregationRules,
	//	pack.AggregationFilters,
	//	aggregationChannel)
	//if err != nil {
	//	return err
	//}
	//
	//// Register connectors
	//parsingChannel := make(chan *sdk.ProcessorTask)
	//registeredConnectors, err := sdk.RegisterConnectors(pack.Connectors, parsingChannel)
	//if err != nil {
	//	return err
	//}
	//
	//// Processor
	//proc := Processor{
	//	InputChannel:     parsingChannel,
	//	OutputChannel:    aggregationChannel,
	//	Workers:          runtime.NumCPU(),
	//	Normalizer:       registeredParsers,
	//	Filters:          registeredFilters,
	//	AggregationRules: registeredAggregationRules,
	//	Connectors:       registeredConnectors,
	//	Destinations:     registeredDestinations,
	//	ID:               collectorID,
	//	Debug:            debugMode,
	//	PrintRaw:         printRaw,
	//	MetricsPort:      metricsPort,
	//}
	//
	//sdk.PrintConfiguration(
	//	registeredConnectors,
	//	registeredParsers,
	//	registeredFilters,
	//	registeredAggregationRules,
	//	registeredDestinations)
	//
	//if err := proc.Start(); err != nil {
	//	return err
	//}
	//
	//runtime.Goexit()
	return nil
}
