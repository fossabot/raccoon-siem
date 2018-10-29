package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk"
)

func Sources(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketSource)
	reply(ctx, err, records)
}

func SourceGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketSource).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("source '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func SourcePUT(ctx *gin.Context) {
	body, err := ctx.GetRawData()

	if err != nil {
		reply(ctx, err)
		return
	}

	s := new(sdk.SourceSettings)
	id, err := unmarshalAndGetID(s, body)

	if err != nil {
		reply(ctx, err)
		return
	}

	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dbBucketSource).Put([]byte(id), body)
	}))
}

func SourceDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketSource).Delete([]byte(id))
	}))
}
