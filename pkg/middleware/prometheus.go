package middleware

import "github.com/prometheus/client_golang/prometheus"

type PrometheusMiddleware struct {
	serviceName     string
	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

func NewPrometheusMiddleware() *PrometheusMiddleware {
	return &PrometheusMiddleware{}
}
