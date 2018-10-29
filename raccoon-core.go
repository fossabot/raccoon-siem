package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/tephrocactus/raccoon-siem/core"
	"log"
)

type coreSettings struct {
	DBFilePath string `long:"db" required:"y" description:"path to database file"`
	Listen     string `long:"listen" default:":7220" description:"http listen address"`
}

func main() {
	// Parse cmd flags
	settings := new(coreSettings)
	if _, err := flags.Parse(settings); err != nil {
		log.Fatal(err)
	}

	// Open database
	core.DBConn = core.NewDB(settings.DBFilePath)

	// Register http endpoints
	httpServer := gin.Default()

	// Parsers
	httpServer.GET("/parsers", core.Parsers)
	httpServer.GET("/parser/:id", core.ParserGET)
	httpServer.PUT("/parser", core.ParserPUT)
	httpServer.DELETE("/parser/:id", core.ParserDELETE)

	// Collectors
	httpServer.GET("/collectors", core.Collectors)
	httpServer.GET("/collector/:id", core.CollectorGET)
	httpServer.PUT("/collector", core.CollectorPUT)
	httpServer.DELETE("/collector/:id", core.CollectorDELETE)

	// Correlators
	httpServer.GET("/correlators", core.Correlators)
	httpServer.GET("/correlator/:id", core.CorrelatorGET)
	httpServer.PUT("/correlator", core.CorrelatorPUT)
	httpServer.DELETE("/correlator/:id", core.CorrelatorDELETE)

	// Correlation Rules
	httpServer.GET("/correlationRules", core.CorrelationRules)
	httpServer.GET("/correlationRule/:id", core.CorrelationRuleGET)
	httpServer.PUT("/correlationRule", core.CorrelationRulePUT)
	httpServer.DELETE("/correlationRule/:id", core.CorrelationRuleDELETE)

	// Aggregation Rules
	httpServer.GET("/aggregationRules", core.AggregationRules)
	httpServer.GET("/aggregationRule/:id", core.AggregationRuleGET)
	httpServer.PUT("/aggregationRule", core.AggregationRulePUT)
	httpServer.DELETE("/aggregationRules/:id", core.AggregationRuleDELETE)

	// Filters
	httpServer.GET("/filters", core.Filters)
	httpServer.GET("/filter/:id", core.FilterGET)
	httpServer.PUT("/filter", core.FilterPUT)
	httpServer.DELETE("/filter/:id", core.FilterDELETE)

	// Active lists
	httpServer.GET("/activeLists", core.ActiveLists)
	httpServer.GET("/activeList/:id", core.ActiveListGET)
	httpServer.PUT("/activeList", core.ActiveListPUT)
	httpServer.DELETE("/activeList/:id", core.ActiveListDELETE)

	// Sources
	httpServer.GET("/sources", core.Sources)
	httpServer.GET("/source/:id", core.SourceGET)
	httpServer.PUT("/source", core.SourcePUT)
	httpServer.DELETE("/source/:id", core.SourceDELETE)

	// Destinations
	httpServer.GET("/destinations", core.Destinations)
	httpServer.GET("/destination/:id", core.DestinationGET)
	httpServer.PUT("/destination", core.DestinationPUT)
	httpServer.DELETE("/destination/:id", core.DestinationDELETE)

	// Dictionaries
	httpServer.GET("/dictionaries", core.Dictionaries)
	httpServer.GET("/dictionary/:id", core.DictionaryGET)
	httpServer.PUT("/dictionary", core.DictionaryPUT)
	httpServer.DELETE("/dictionary/:id", core.DictionaryDELETE)

	// Component registration
	httpServer.GET("/register/collector/:id", core.CollectorRegister)
	httpServer.GET("/register/correlator/:id", core.CorrelatorRegister)

	// Run http server
	httpServer.Run(settings.Listen)
}
