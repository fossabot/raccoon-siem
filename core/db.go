package core

import (
	"github.com/boltdb/bolt"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"gopkg.in/yaml.v2"
	"time"
)

var (
	DBConn                  *DB
	dbBucketCorrelationRule = []byte("correlationRule")
	dbBucketAggregationRule = []byte("aggregationRule")
	dbBucketFilter          = []byte("filter")
	dbBucketParser          = []byte("parser")
	dbBucketCollector       = []byte("collector")
	dbBucketCorrelator      = []byte("correlator")
	dbBucketActiveList      = []byte("activeList")
	dbBucketConnector       = []byte("connector")
	dbBucketDestination     = []byte("destination")
	dbBucketDictionary      = []byte("dictionary")
)

var bucketNames = [][]byte{
	dbBucketCorrelationRule, dbBucketFilter, dbBucketParser, dbBucketCollector,
	dbBucketCorrelator, dbBucketActiveList, dbBucketConnector, dbBucketAggregationRule,
	dbBucketDestination, dbBucketDictionary,
}

func NewDB(path string) *DB {
	db := new(DB)

	descriptor, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 5 * time.Second})
	sdk.PanicOnError(err)

	db.h = descriptor

	err = db.createBuckets()
	sdk.PanicOnError(err)

	return db
}

type DB struct {
	h *bolt.DB
}

func (db *DB) ListKeys(bucket []byte) ([]byte, error) {
	result := make([]string, 0)

	err := db.h.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucket).Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			result = append(result, string(k))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return yaml.Marshal(result)
}

func (db *DB) createBuckets() error {
	return db.h.Update(func(tx *bolt.Tx) error {
		for _, name := range bucketNames {
			_, err := tx.CreateBucketIfNotExists(name)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
