package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk/aggregation"
)

func AggregationRules(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketAggregationRule)
	reply(ctx, err, records)
}

func AggregationRuleGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketAggregationRule).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("aggregation rule '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func AggregationRulePUT(ctx *gin.Context) {
	body, err := ctx.GetRawData()

	if err != nil {
		reply(ctx, err)
		return
	}

	s := new(aggregation.Config)
	id, err := unmarshalAndGetID(s, body)

	if err != nil {
		reply(ctx, err)
		return
	}

	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dbBucketAggregationRule).Put([]byte(id), body)
	}))
}

func AggregationRuleDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketAggregationRule).Delete([]byte(id))
	}))
}
