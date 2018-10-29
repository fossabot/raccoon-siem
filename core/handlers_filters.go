package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/tephrocactus/raccoon-siem/sdk"
)

func Filters(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketFilter)
	reply(ctx, err, records)
}

func FilterGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketFilter).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("filter '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func FilterPUT(ctx *gin.Context) {
	body, err := ctx.GetRawData()

	if err != nil {
		reply(ctx, err)
		return
	}

	s := new(sdk.FilterSettings)
	id, err := unmarshalAndGetID(s, body)

	if err != nil {
		reply(ctx, err)
		return
	}

	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(dbBucketFilter).Put([]byte(id), body)
	}))
}

func FilterDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketFilter).Delete([]byte(id))
	}))
}
