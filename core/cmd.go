package core

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tephrocactus/raccoon-siem/core/db"
	"github.com/tephrocactus/raccoon-siem/core/migrator"
)

var (
	Cmd = &cobra.Command{
		Use:   "core",
		Short: "start raccoon configuration server",
		Args:  cobra.ExactArgs(0),
		RunE:  run,
	}

	// String flags variables
	listen, dbHost, dbPort, dbScheme string
)

func init() {
	// Migration
	Cmd.AddCommand(migrator.Cmd)

	// Cockroach
	Cmd.PersistentFlags().StringVar(
		&dbHost,
		"db.host",
		"localhost",
		"database host",
	)
	Cmd.PersistentFlags().StringVar(
		&dbPort,
		"db.port",
		"26257",
		"database port",
	)
	Cmd.PersistentFlags().StringVarP(
		&dbScheme,
		"db.scheme",
		"d",
		"raccoon",
		"database scheme",
	)

	// Listen address
	Cmd.Flags().StringVarP(
		&listen,
		"listen",
		"l",
		":7220",
		"listen address")

}

func run(_ *cobra.Command, _ []string) error {
	// Open database
	var err error
	DBConn, err = db.Connect(dbHost, dbPort, dbScheme)
	if err != nil {
		return err
	}

	// Register http endpoints
	router := gin.Default()

	config := router.Group("/config")

	//// normalizer
	//router.GET("/parser", Normalizers)
	//router.GET("/parser/:id", NormalizerGET)
	//router.PUT("/parser", NormalizerPUT)
	//router.DELETE("/parser/:id", NormalizerDELETE)
	//
	//// Collectors
	//router.GET("/collector", Collectors)
	//router.GET("/collector/:id", CollectorGET)
	//router.PUT("/collector", CollectorPUT)
	//router.DELETE("/collector/:id", CollectorDELETE)
	//
	//// Correlators
	//router.GET("/correlator", Correlators)
	//router.GET("/correlator/:id", CorrelatorGET)
	//router.PUT("/correlator", CorrelatorPUT)
	//router.DELETE("/correlator/:id", CorrelatorDELETE)
	//
	//// Correlation Rules
	//router.GET("/correlationRule", CorrelationRules)
	//router.GET("/correlationRule/:id", CorrelationRuleGET)
	//router.PUT("/correlationRule", CorrelationRulePUT)
	//router.DELETE("/correlationRule/:id", CorrelationRuleDELETE)
	//
	//// Aggregation Rules
	//router.GET("/aggregationRule", AggregationRules)
	//router.GET("/aggregationRule/:id", AggregationRuleGET)
	//router.PUT("/aggregationRule", AggregationRulePUT)
	//router.DELETE("/aggregationRule/:id", AggregationRuleDELETE)
	//
	//// Filters
	//router.GET("/filter", Filters)
	//router.GET("/filter/:id", FilterGET)
	//router.PUT("/filter", FilterPUT)
	//router.DELETE("/filter/:id", FilterDELETE)
	//
	//// Active lists
	//router.GET("/activeList", ActiveLists)
	//router.GET("/activeList/:id", ActiveListGET)
	//router.PUT("/activeList", ActiveListPUT)
	//router.DELETE("/activeList/:id", ActiveListDELETE)
	//
	//// Connectors
	//router.GET("/connector", ConnectorsList)
	//router.GET("/connector/:id", ConnectorGET)
	//router.PUT("/connector", ConnectorPUT)
	//router.DELETE("/connector/:id", ConnectorDELETE)
	//
	// destinations
	destinationGroup := config.Group("/destination")
	destinationGroup.GET("/", Destinations)
	destinationGroup.GET("/destination/:id", DestinationGET)
	destinationGroup.PUT("/destination", DestinationPUT)
	destinationGroup.DELETE("/destination/:id", DestinationDELETE)
	//
	//// Dictionaries
	//router.GET("/dictionary", Dictionaries)
	//router.GET("/dictionary/:id", DictionaryGET)
	//router.PUT("/dictionary", DictionaryPUT)
	//router.DELETE("/dictionary/:id", DictionaryDELETE)
	//
	//// Component registration
	//router.GET("/register/collector/:id", CollectorRegister)
	//router.GET("/register/correlator/:id", CorrelatorRegister)

	// Run http server
	return router.Run(listen)
}
