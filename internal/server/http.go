package server

import (
	"kratosent/internal/conf"
	"kratosent/internal/service"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
	registers []service.RegisterHTTPFunc,
	histoVec *prometheus.HistogramVec,
	counterVec *prometheus.CounterVec,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(histoVec)),
				metrics.WithRequests(prom.NewCounter(counterVec)),
			),
			mmd.Server(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)

	for _, fc := range registers {
		fc(srv)
	}

	prometheus.MustRegister(histoVec, counterVec)
	srv.Handle("/metrics", promhttp.Handler())

	return srv
}
