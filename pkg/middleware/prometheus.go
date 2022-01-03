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

var BOOK_RESERVED = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "count_of_books_reserved",
	Help: "The total number of processed events",
}, []string{"book_id", "status_code"})

func NewPrometheusMiddleware(sName string) *PrometheusMiddleware {
	requestCount := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "reserve_task",
		Help: "The total number of processed events",
	}, []string{"method", "path", "status_code", "date"})
	requestDuration := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "hist_creator",
		Help:    "Histogram data",
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
			"status_code": strconv.Itoa(c.Writer.Status()),
			"date":        time.Now().Format("01-02-2006")}).Inc()
		go p.requestDuration.With(prometheus.Labels{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"status_code": strconv.Itoa(c.Writer.Status())}).Observe(float64(time.Since(begin)) / float64(time.Second))
	}
}
