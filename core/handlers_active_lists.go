package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func ActiveLists(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketActiveList)
	reply(ctx, err, records)
}

func ActiveListGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketActiveList).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("active list '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func ActiveListPUT(ctx *gin.Context) {
	//body, err := ctx.GetRawData()
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//s := new(sdk.ActiveListSettings)
	//id, err := unmarshalAndGetID(s, body)
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
	//	return tx.Bucket(dbBucketActiveList).Put([]byte(id), body)
	//}))
}

func ActiveListDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketActiveList).Delete([]byte(id))
	}))
}
