package core

import (
	"github.com/gin-gonic/gin"
)

func CollectorRegister(ctx *gin.Context) {
	//id := ctx.Param("id")
	//pack := sdk.CollectorPackage{}
	//
	//err := DBConn.h.View(func(tx *bolt.Tx) (err error) {
	//	collectorsBucket := tx.Bucket(dbBucketCollector)
	//	data := collectorsBucket.Get([]byte(id))
	//
	//	if data == nil {
	//		return fmt.Errorf("collector '%s' does not exist", id)
	//	}
	//
	//	collectorSettings := new(sdk.CollectorConfig)
	//
	//	if err = yaml.Unmarshal(data, collectorSettings); err != nil {
	//		return err
	//	}
	//
	//	pack.Connectors, err = readConnectorsByIDs(collectorSettings.Connectors, tx)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	pack.Normalizers, err = readNormalizersByIDs(collectorSettings.Parsers, true, nil, tx)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	pack.Destinations, err = readDestinationsByIDs(collectorSettings.Destinations, tx)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	pack.AggregationRules, err = readAggregationRulesByIDs(collectorSettings.AggregationRules, tx)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	aggregationFiltersIDs := make([]string, 0)
	//
	//	for _, rule := range pack.AggregationRules {
	//		for _, evt := range rule.Events {
	//			aggregationFiltersIDs = append(aggregationFiltersIDs, evt.Filter)
	//		}
	//	}
	//
	//	pack.AggregationFilters, err = readFiltersByIDs(aggregationFiltersIDs, nil, tx)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	filterIDs := make([]string, 0)
	//
	//	for _, fName := range collectorSettings.Filters {
	//		filterIDs = append(filterIDs, fName)
	//	}
	//
	//	pack.Filters, err = readFiltersByIDs(filterIDs, nil, tx)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	pack.Dictionaries, err = readAllDictionaries(tx)
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	pack.Debug = collectorSettings.Debug
	//	pack.Name = collectorSettings.Name
	//	pack.MeterPeriod = collectorSettings.MeterPeriod
	//
	//	return nil
	//})
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//replyData, err := yaml.Marshal(pack)
	//reply(ctx, err, replyData)
}
