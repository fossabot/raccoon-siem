package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk"
)

func Correlators(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketCorrelator)
	reply(ctx, err, records)
}

func CorrelatorGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketCorrelator).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("correlator '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func CorrelatorPUT(ctx *gin.Context) {
	body, err := ctx.GetRawData()

	if err != nil {
		reply(ctx, err)
		return
	}

	s := new(sdk.CorrelatorConfig)
	id, err := unmarshalAndGetID(s, body)

	if err != nil {
		reply(ctx, err)
		return
	}

	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dbBucketCorrelator).Put([]byte(id), body)
	}))
}

func CorrelatorDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketCorrelator).Delete([]byte(id))
	}))
}
