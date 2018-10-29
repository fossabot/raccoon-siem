package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk"
)

func Parsers(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketParser)
	reply(ctx, err, records)
}

func ParserGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketParser).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("parser '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func ParserPUT(ctx *gin.Context) {
	body, err := ctx.GetRawData()

	if err != nil {
		reply(ctx, err)
		return
	}

	s := new(sdk.ParserSettings)
	id, err := unmarshalAndGetID(s, body)

	if err != nil {
		reply(ctx, err)
		return
	}

	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dbBucketParser).Put([]byte(id), body)
	}))
}

func ParserDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketParser).Delete([]byte(id))
	}))
}
