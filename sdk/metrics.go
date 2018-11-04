package sdk

import "github.com/prometheus/client_golang/prometheus"

const (
	MetricsNamespace           = "raccoon"
	MetricsSubsystemCollector  = "collector"
	MetricsSubsystemCorrelator = "correlator"
)

func MetricsDefaultLatencyBuckets() []float64 {
	return prometheus.ExponentialBuckets(100, 10, 8)
}
