package start

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of incoming HTTP requests",
		},
		[]string{"method", "endpoint"},
	)

	httpRequestsInProgress = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_progress",
			Help: "Current number of HTTP requests in progress",
		},
		[]string{"method", "endpoint"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of completed HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint", "status_code"},
	)
)

func init() {
	// Register metrics with prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestsInProgress)
	prometheus.MustRegister(httpRequestDuration)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpRequestsInProgress.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		startTime := time.Now()

		c.Next()

		httpRequestsInProgress.WithLabelValues(c.Request.Method, c.FullPath()).Dec()
		httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, c.Request.URL.Path, strconv.Itoa(c.Writer.Status())).Observe(time.Since(startTime).Seconds())
	}
}
