package cmd

import (
	"fmt"
	"github.com/getsentry/sentry-go"
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

		defer func() {
			// If something panics, recover and report to Sentry.
			if err := recover(); err != nil {
				sentry.CaptureException(fmt.Errorf("panic in PrometheusMiddleware: %v", err))
				sentry.Flush(time.Second * 5)
			}

			// Decrement requests in progress and increment total requests counters.
			httpRequestsInProgress.WithLabelValues(c.Request.Method, c.FullPath()).Dec()
			httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath()).Inc()

			// Record the duration of the request.
			httpRequestDuration.WithLabelValues(c.Request.Method, c.Request.URL.Path, strconv.Itoa(c.Writer.Status())).Observe(time.Since(startTime).Seconds())
		}()

		// Continue with the next middleware/handler function.
		c.Next()
	}
}
