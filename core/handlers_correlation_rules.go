package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk/correlation"
)

func CorrelationRules(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketCorrelationRule)
	reply(ctx, err, records)
}

func CorrelationRuleGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketCorrelationRule).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("correlation rule '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func CorrelationRulePUT(ctx *gin.Context) {
	body, err := ctx.GetRawData()

	if err != nil {
		reply(ctx, err)
		return
	}

	s := new(correlation.Config)
	id, err := unmarshalAndGetID(s, body)

	if err != nil {
		reply(ctx, err)
		return
	}

	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dbBucketCorrelationRule).Put([]byte(id), body)
	}))
}

func CorrelationRuleDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketCorrelationRule).Delete([]byte(id))
	}))
}
