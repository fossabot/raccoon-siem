package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"gopkg.in/yaml.v2"
)

func CorrelatorRegister(ctx *gin.Context) {
	id := ctx.Param("id")
	pack := sdk.CorrelatorPackage{}

	err := DBConn.h.View(func(tx *bolt.Tx) (err error) {
		correlatorsBucket := tx.Bucket(dbBucketCorrelator)
		data := correlatorsBucket.Get([]byte(id))

		if data == nil {
			return fmt.Errorf("correlator '%s' does not exist", id)
		}

		correlatorSettings := new(sdk.CorrelatorSettings)

		if err = yaml.Unmarshal(data, correlatorSettings); err != nil {
			return err
		}

		pack.Connectors, err = readConnectorsByIDs(correlatorSettings.Connectors, tx)

		if err != nil {
			return err
		}

		pack.Destinations, err = readDestinationsByIDs(correlatorSettings.Destinations, tx)

		if err != nil {
			return err
		}

		pack.CorrelationRules, err = readCorrelationRulesByIDs(correlatorSettings.CorrelationRules, tx)

		if err != nil {
			return err
		}

		filterIDs := make([]string, 0)

		for _, rule := range pack.CorrelationRules {
			if rule.Filter != "" {
				filterIDs = append(filterIDs, rule.Filter)
			}

			for _, evt := range rule.Events {
				filterIDs = append(filterIDs, evt.Filter)
			}
		}

		pack.Filters, err = readFiltersByIDs(filterIDs, nil, tx)

		if err != nil {
			return err
		}

		pack.ActiveLists, err = readAllActiveLists(tx)

		if err != nil {
			return err
		}

		pack.Debug = correlatorSettings.Debug
		pack.Name = correlatorSettings.Name
		pack.MeterPeriod = correlatorSettings.MeterPeriod

		return nil
	})

	if err != nil {
		reply(ctx, err)
		return
	}

	replyData, err := yaml.Marshal(pack)
	reply(ctx, err, replyData)
}
