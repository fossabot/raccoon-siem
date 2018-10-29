package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk"
)

func Destinations(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketDestination)
	reply(ctx, err, records)
}

func DestinationGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketDestination).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("destination '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func DestinationPUT(ctx *gin.Context) {
	body, err := ctx.GetRawData()

	if err != nil {
		reply(ctx, err)
		return
	}

	s := new(sdk.DestinationSettings)
	id, err := unmarshalAndGetID(s, body)

	if err != nil {
		reply(ctx, err)
		return
	}

	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dbBucketDestination).Put([]byte(id), body)
	}))
}

func DestinationDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketDestination).Delete([]byte(id))
	}))
}
