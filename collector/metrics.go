package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"log"
	"net/http"
	"time"
)

type metrics struct {
	events               *prometheus.CounterVec
	filtered             *prometheus.CounterVec
	aggregated           *prometheus.CounterVec
	latencyOverall       prometheus.Histogram
	latencyParsing       *prometheus.HistogramVec
	latencySerialization *prometheus.HistogramVec
	port                 string
}

func (r *metrics) runServer() *metrics {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":"+r.port, nil))
	}()
	return r
}

func (r *metrics) registerEventInput(connector string) {
	r.events.WithLabelValues("in", connector).Inc()
}

func (r *metrics) registerEventOutput(connector string) {
	r.events.WithLabelValues("out", connector).Inc()
}

func (r *metrics) registerEventFiltration(filter string, connector string) {
	r.filtered.WithLabelValues(filter, connector).Inc()
}

func (r *metrics) registerEventAggregation(rule string, count int, connector string) {
	r.aggregated.WithLabelValues(rule, connector).Add(float64(count))
}

func (r *metrics) registerParsingDuration(parser string, start time.Time) {
	r.latencyParsing.
		WithLabelValues(parser).
		Observe(float64(time.Since(start).Nanoseconds()))
}

func (r *metrics) registerSerializationDuration(destination string, start time.Time) {
	r.latencySerialization.
		WithLabelValues(destination).
		Observe(float64(time.Since(start).Nanoseconds()))
}

func (r *metrics) registerOverallProcessingDuration(start time.Time) {
	r.latencyOverall.Observe(float64(time.Since(start).Nanoseconds()))
}

func newMetrics(port string) *metrics {
	m := &metrics{
		port: port,
		events: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "events",
			}, []string{"direction", "source"}),
		filtered: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "events_filtered",
			}, []string{"filter", "source"}),
		aggregated: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "events_aggregated",
			}, []string{"rule", "source"}),
		latencyOverall: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "latency_total",
				Buckets:   sdk.MetricsDefaultLatencyBuckets(),
			}),
		latencyParsing: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "latency_parsing",
				Buckets:   sdk.MetricsDefaultLatencyBuckets(),
			}, []string{"parser"}),
		latencySerialization: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "latency_serialization",
				Buckets:   sdk.MetricsDefaultLatencyBuckets(),
			}, []string{"destination"}),
	}

	prometheus.MustRegister(m.events)
	prometheus.MustRegister(m.filtered)
	prometheus.MustRegister(m.aggregated)
	prometheus.MustRegister(m.latencyOverall)
	prometheus.MustRegister(m.latencyParsing)
	prometheus.MustRegister(m.latencySerialization)

	return m
}
