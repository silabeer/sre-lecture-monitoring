package middleware

import (

	"demomon/internal/metrics"
	"github.com/gin-gonic/gin"
)

func InflightRequests() gin.HandlerFunc {
	// Return a new middleware handler function
	return func(c *gin.Context) {
		metrics.HttpRequestsCurrent.WithLabelValues().Inc()
		// Process the request
		c.Next()
		metrics.HttpRequestsCurrent.WithLabelValues().Dec()
	}
}