package core

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "core",
		Short: "start raccoon configuration server",
		Args:  cobra.ExactArgs(0),
		RunE:  run,
	}

	// String flags variables
	listen, dbFile string
)

func init() {
	// Listen address
	Cmd.Flags().StringVarP(
		&listen,
		"listen",
		"l",
		":7220",
		"listen address")

	// Raccoon core DB file
	Cmd.Flags().StringVarP(
		&dbFile,
		"db",
		"d",
		"raccoon.db",
		"database file")
}

func run(_ *cobra.Command, _ []string) error {
	// Open database
	DBConn = NewDB(dbFile)

	// Register http endpoints
	httpServer := gin.Default()

	// Parsers
	httpServer.GET("/parser", Parsers)
	httpServer.GET("/parser/:id", ParserGET)
	httpServer.PUT("/parser", ParserPUT)
	httpServer.DELETE("/parser/:id", ParserDELETE)

	// Collectors
	httpServer.GET("/collector", Collectors)
	httpServer.GET("/collector/:id", CollectorGET)
	httpServer.PUT("/collector", CollectorPUT)
	httpServer.DELETE("/collector/:id", CollectorDELETE)

	// Correlators
	httpServer.GET("/correlator", Correlators)
	httpServer.GET("/correlator/:id", CorrelatorGET)
	httpServer.PUT("/correlator", CorrelatorPUT)
	httpServer.DELETE("/correlator/:id", CorrelatorDELETE)

	// Correlation Rules
	httpServer.GET("/correlationRule", CorrelationRules)
	httpServer.GET("/correlationRule/:id", CorrelationRuleGET)
	httpServer.PUT("/correlationRule", CorrelationRulePUT)
	httpServer.DELETE("/correlationRule/:id", CorrelationRuleDELETE)

	// Aggregation Rules
	httpServer.GET("/aggregationRule", AggregationRules)
	httpServer.GET("/aggregationRule/:id", AggregationRuleGET)
	httpServer.PUT("/aggregationRule", AggregationRulePUT)
	httpServer.DELETE("/aggregationRule/:id", AggregationRuleDELETE)

	// Filters
	httpServer.GET("/filter", Filters)
	httpServer.GET("/filter/:id", FilterGET)
	httpServer.PUT("/filter", FilterPUT)
	httpServer.DELETE("/filter/:id", FilterDELETE)

	// Active lists
	httpServer.GET("/activeList", ActiveLists)
	httpServer.GET("/activeList/:id", ActiveListGET)
	httpServer.PUT("/activeList", ActiveListPUT)
	httpServer.DELETE("/activeList/:id", ActiveListDELETE)

	// Connectors
	httpServer.GET("/connector", ConnectorsList)
	httpServer.GET("/connector/:id", ConnectorGET)
	httpServer.PUT("/connector", ConnectorPUT)
	httpServer.DELETE("/connector/:id", ConnectorDELETE)

	// Destinations
	httpServer.GET("/destination", Destinations)
	httpServer.GET("/destination/:id", DestinationGET)
	httpServer.PUT("/destination", DestinationPUT)
	httpServer.DELETE("/destination/:id", DestinationDELETE)

	// Dictionaries
	httpServer.GET("/dictionary", Dictionaries)
	httpServer.GET("/dictionary/:id", DictionaryGET)
	httpServer.PUT("/dictionary", DictionaryPUT)
	httpServer.DELETE("/dictionary/:id", DictionaryDELETE)

	// Component registration
	httpServer.GET("/register/collector/:id", CollectorRegister)
	httpServer.GET("/register/correlator/:id", CorrelatorRegister)

	// Storage mapping generator
	httpServer.GET("/storage/template", GenerateStorageMapping)

	// Run http server
	return httpServer.Run(listen)
}
