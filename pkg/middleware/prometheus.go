package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
)

type PrometheusMiddleware struct {
	serviceName     string
	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

func NewPrometheusMiddleware(sName string) *PrometheusMiddleware {
	requestCount := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "request_of_method",
		Help: "The total number of processed events",
	}, []string{"method", "path", "statuscode"})
	requestDuration := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "hyst_creator",
		Help:    "Hystogram data",
		Buckets: []float64{1, 2, 5, 10, 20, 60},
	}, []string{})
	return &PrometheusMiddleware{
		serviceName:     sName,
		requestCount:    requestCount,
		requestDuration: requestDuration,
	}
}

func (p *PrometheusMiddleware) Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.requestCount.With(prometheus.Labels{
			"method":     c.HandlerName(),
			"path":       c.Request.RequestURI,
			"statuscode": strconv.Itoa(c.Writer.Status())}).Inc()
		c.Next()
	}
}
