package api

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(txMiddleware())

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

	//
	//
	//// Component registration
	//router.GET("/register/collector/:id", CollectorRegister)
	//router.GET("/register/correlator/:id", CorrelatorRegister)

	// Connectors
	connectorGroup := config.Group("/connector")
	connectorGroup.GET("/", readConnectors)
	connectorGroup.GET("/:id", readConnector)
	connectorGroup.POST("/", createConnector)
	connectorGroup.PUT("/:id", updateConnector)
	connectorGroup.DELETE("/:id", deleteConnector)

	// Dictionaries
	dictionaryGroup := config.Group("/dictionary")
	dictionaryGroup.GET("/", readDictionaries)
	dictionaryGroup.GET("/:id", readDictionary)
	dictionaryGroup.POST("/", createDictionary)
	dictionaryGroup.PUT("/:id", updateDictionary)
	dictionaryGroup.DELETE("/:id", deleteDictionary)

	// Destination configs
	destinationGroup := config.Group("/destination")
	destinationGroup.GET("/", readDestinations)
	destinationGroup.POST("/", createDestination)
	destinationGroup.GET("/:id", readDestination)
	destinationGroup.PUT("/:id", updateDestination)
	destinationGroup.DELETE("/:id", deleteDestination)

	return router
}
