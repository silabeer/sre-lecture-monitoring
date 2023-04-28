package middleware

import (
 
    "time"
	"strconv"
	"github.com/gin-gonic/gin"
	"demomon/internal/metrics"
)

func HTTPMetrics() gin.HandlerFunc {
	// Return a new middleware handler function
	return func(c *gin.Context) {
		// Start a timer to measure the request duration
		start := time.Now()
		// Process the request
		c.Next()
		// Record the request duration
		duration := time.Since(start).Seconds()
		status := c.Writer.Status()
		path := c.Request.URL.Path
		method := c.Request.Method
		metrics.HttpRequestsTotal.WithLabelValues(path, method,strconv.Itoa(status)).Inc()
		metrics.HttpRequestsDurationHistorgram.WithLabelValues(path, method).Observe(duration)
		metrics.HttpRequestsDurationSummary.WithLabelValues(path, method).Observe(duration)
	}
}