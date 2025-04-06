package metrics

import (
	"flag"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
)

var (
	Namespace = flag.String("namespace", "web", "metrics namespace")
	Subsystem = flag.String("subsystem", "server1", "metrics subsystem")

	Requests metrics.Counter = prometheus.NewCounterFrom(
		prom.CounterOpts{
			Namespace: *Namespace,
			Subsystem: *Subsystem,
			Name:      "request_count",
			Help:      "Total requests",
		},
		[]string{},
	)

	writeErrors metrics.Counter = prometheus.NewCounterFrom(
		prom.CounterOpts{
			Namespace: *Namespace,
			Subsystem: *Subsystem,
			Name:      "write_error_count",
			Help:      "Total write errors",
		},
		[]string{},
	)

	OpenConnections metrics.Gauge = prometheus.NewGaugeFrom(
		prom.GaugeOpts{
			Namespace: *Namespace,
			Subsystem: *Subsystem,
			Name:      "open_connections",
			Help:      "current open connections",
		},
		[]string{},
	)

	RequestDuration metrics.Histogram = prometheus.NewHistogramFrom(
		prom.HistogramOpts{
			Namespace: *Namespace,
			Subsystem: *Subsystem,
			Buckets: []float64{
				0.000_000_1, 0.000_000_2, 0.000_000_3, 0.000_000_4, 0.000_000_5,
				0.000_001, 0.000_002_5, 0.000_005, 0.000_007_5, 0.000_01, 0.000_1, 0.001, 0.01,
			},
			Name: "request_duration_histogram_seconds",
			Help: "Total duration of all requests",
		},
		[]string{},
	)
)
