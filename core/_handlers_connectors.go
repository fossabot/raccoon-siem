package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

func ConnectorsList(ctx *gin.Context) {
	records, err := DBConn.ListKeys(dbBucketConnector)
	reply(ctx, err, records)
}

func ConnectorGET(ctx *gin.Context) {
	var replyData []byte

	err := DBConn.h.View(func(tx *bolt.Tx) error {
		id := ctx.Param("id")

		value := tx.Bucket(dbBucketConnector).Get([]byte(id))

		if value == nil {
			return fmt.Errorf("connector '%s' does not exist", id)
		}

		replyData = value
		return nil
	})

	reply(ctx, err, replyData)
}

func ConnectorPUT(ctx *gin.Context) {
	//body, err := ctx.GetRawData()
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//s := new(connectors.Config)
	//id, err := unmarshalAndGetID(s, body)
	//
	//if err != nil {
	//	reply(ctx, err)
	//	return
	//}
	//
	//reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
	//	return tx.Bucket(dbBucketConnector).Put([]byte(id), body)
	//}))
}

func ConnectorDELETE(ctx *gin.Context) {
	reply(ctx, DBConn.h.Update(func(tx *bolt.Tx) error {
		id := ctx.Param("id")
		return tx.Bucket(dbBucketConnector).Delete([]byte(id))
	}))
}
