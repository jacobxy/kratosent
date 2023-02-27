package middleware

import (
	// prometheusv2 "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"kratosent/internal/conf"

	prometheus "github.com/prometheus/client_golang/prometheus"
)

func NewPrometheusHistogramVec(promConfig *conf.Middleware) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: promConfig.Prometheus.Histogram.Namespace,
		Subsystem: promConfig.Prometheus.Histogram.Subsystem,
		Name:      promConfig.Prometheus.Histogram.Name,
		Help:      "server requests duration seconds",
		Buckets:   []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.250, 0.5, 1},
	}, []string{"kind", "operation"})
}

func NewPrometheusCounterVec(promConfig *conf.Middleware) *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: promConfig.Prometheus.Counter.Namespace,
		Subsystem: promConfig.Prometheus.Counter.Subsystem,
		Name:      promConfig.Prometheus.Counter.Name,
		Help:      "The total number of processed requests",
	}, []string{"kind", "operation", "code", "reason"})
}
