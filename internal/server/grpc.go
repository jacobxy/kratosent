package server

import (
	"kratosent/internal/conf"
	"kratosent/internal/service"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/prometheus/client_golang/prometheus"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server,
	registers []service.RegisterGRPCFunc,
	histoVec *prometheus.HistogramVec,
	counterVec *prometheus.CounterVec,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(histoVec)),
				metrics.WithRequests(prom.NewCounter(counterVec)),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	for _, fc := range registers {
		fc(srv)
	}

	return srv
}
