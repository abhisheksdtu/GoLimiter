package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rate_limiter_requests_total",
		Help: "Total number of rate limiter requests.",
	})
	requestsAllowed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rate_limiter_allowed_total",
		Help: "Total number of allowed requests.",
	})
	requestsDenied = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rate_limiter_denied_total",
		Help: "Total number of denied requests.",
	})
)

// PrometheusMiddleware returns a Gin middleware for Prometheus metrics.
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestsTotal.Inc()
		c.Next()
	}
}

// MetricsHandler returns a Gin handler for serving Prometheus metrics.
func MetricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// RecordMetrics records the outcome of a rate limit check.
func RecordMetrics(allowed bool) {
	if allowed {
		requestsAllowed.Inc()
	} else {
		requestsDenied.Inc()
	}
}
