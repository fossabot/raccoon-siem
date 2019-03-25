package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func Dictionaries(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketDictionary)
	reply(ctx, err, records)
}

func DictionaryGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketDictionary).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("dictionary '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func DictionaryPUT(ctx *gin.Context) {
	//body, err := ctx.GetRawData()
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//s := new(sdk.DictionarySettings)
	//id, err := unmarshalAndGetID(s, body)
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
	//	return tx.Bucket(dbBucketDictionary).Put([]byte(id), body)
	//}))
}

func DictionaryDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketDictionary).Delete([]byte(id))
	}))
}
