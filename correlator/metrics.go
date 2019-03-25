package correlator

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/tephrocactus/raccoon-siem/sdk"
	"net/http"
	"time"
)

type metrics struct {
	eventsReceived    prometheus.Counter
	eventsCorrelated  *prometheus.CounterVec
	eventsSent        *prometheus.CounterVec
	eventsProcessed   prometheus.Counter
	processingLatency prometheus.Histogram
	port              string
}

func (r *metrics) eventReceived() {
	r.eventsReceived.Inc()
}

func (r *metrics) eventSent(destination string) {
	r.eventsSent.WithLabelValues(destination).Inc()
}

func (r *metrics) eventCorrelated(rule string) {
	r.eventsCorrelated.WithLabelValues(rule).Inc()
}

func (r *metrics) eventProcessed() {
	r.eventsProcessed.Inc()
}

func (r *metrics) processingTook(took time.Duration) {
	r.processingLatency.Observe(float64(took.Nanoseconds()))
}

func newMetrics(port string) *metrics {
	m := &metrics{
		port: port,
		eventsReceived: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "eventsReceived",
			}),
		eventsProcessed: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "eventsProcessed",
			}),
		eventsCorrelated: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "eventsCorrelated",
			}, []string{"rule"}),
		eventsSent: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "eventsSent",
			}, []string{"destination"}),
		processingLatency: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: sdk.MetricsNamespace,
				Subsystem: sdk.MetricsSubsystemCollector,
				Name:      "processingLatency",
				Buckets:   prometheus.DefBuckets,
			}),
	}

	prometheus.MustRegister(
		m.eventsReceived,
		m.eventsProcessed,
		m.eventsCorrelated,
		m.eventsSent,
		m.processingLatency)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":"+m.port, nil))
	}()

	return m
}
