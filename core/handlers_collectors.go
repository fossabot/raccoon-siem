package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func Collectors(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketCollector)
	reply(ctx, err, records)
}

func CollectorGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketCollector).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("collector '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func CollectorPUT(ctx *gin.Context) {
	//body, err := ctx.GetRawData()
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//s := new(sdk.CollectorConfig)
	//id, err := unmarshalAndGetID(s, body)
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
	//	return tx.Bucket(dbBucketCollector).Put([]byte(id), body)
	//}))
}

func CollectorDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketCollector).Delete([]byte(id))
	}))
}
