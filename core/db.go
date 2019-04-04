package core

import (
	"database/sql"
	"github.com/tephrocactus/raccoon-siem/core/db"
	_ "github.com/tephrocactus/raccoon-siem/core/migrator"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	DBConn                  *sql.DB
	UDBConn                 sqlbuilder.Database
	dbBucketCorrelationRule = []byte("correlationRule")
	dbBucketAggregationRule = []byte("aggregationRule")
	dbBucketFilter          = []byte("filter")
	dbBucketNormalizer      = []byte("normalizer")
	dbBucketCollector       = []byte("collector")
	dbBucketCorrelator      = []byte("correlator")
	dbBucketActiveList      = []byte("activeList")
	dbBucketConnector       = []byte("connector")
	dbBucketDestination     = []byte("destination")
	dbBucketDictionary      = []byte("dictionary")
)

func NewUdbConnection() error {
	var err error
	UDBConn, err = db.ConnectUdb(dbHost, dbPort, dbScheme)
	if err != nil {
		return err
	}

	return nil
}