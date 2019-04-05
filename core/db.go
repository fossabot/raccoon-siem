package core

import (
	_ "github.com/tephrocactus/raccoon-siem/core/migrator"
)

var (
	dbBucketCorrelationRule = []byte("correlationRule")
	dbBucketAggregationRule = []byte("aggregationRule")
	dbBucketFilter          = []byte("filter")
	dbBucketNormalizer      = []byte("normalizer")
	dbBucketCollector       = []byte("collector")
	dbBucketCorrelator      = []byte("correlator")
	dbBucketActiveList      = []byte("activeList")
	dbBucketConnector       = []byte("connector")
)
