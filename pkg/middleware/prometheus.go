package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"time"
)

type PrometheusMiddleware struct {
	serviceName     string
	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

func NewPrometheusMiddleware(sName string) *PrometheusMiddleware {
	requestCount := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "reserve_task",
		Help: "The total number of processed events",
	}, []string{"method", "path", "status_code"})
	requestDuration := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "hyst_creator",
		Help:    "Hystogram data",
		Buckets: []float64{.05, .1, 1, 2, 5, 15},
	}, []string{"method", "path", "status_code"})
	return &PrometheusMiddleware{
		serviceName:     sName,
		requestCount:    requestCount,
		requestDuration: requestDuration,
	}
}

func (p *PrometheusMiddleware) Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		c.Next()
		go p.requestCount.With(prometheus.Labels{
			"method":      c.Request.Method,
			"path":        c.Request.RequestURI,
			"status_code": strconv.Itoa(c.Writer.Status())}).Inc()
		go p.requestDuration.WithLabelValues(
			(c.Request.Method),
			c.Request.URL.Path,
			strconv.Itoa(c.Writer.Status()),
		).Observe(float64(time.Since(begin)) / float64(time.Second))
	}
}
