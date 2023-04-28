package main

import (
	"demomon/internal/handlers"
	"demomon/internal/metrics"
	"demomon/internal/middleware"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getHTTPRequestsInflightMax() float64 {
	httpRequestsInflightMaxString := os.Getenv("HTTP_REQUESTS_INFLIGHT_MAX")
	httpRequestsInflightMax := 20.0
	if httpRequestsInflightMaxString != "" {
		httpRequestsInflightMax, _ = strconv.ParseFloat(httpRequestsInflightMaxString, 32)
	}

	return httpRequestsInflightMax
}

func main() {

	httpRequestsInflightMax := getHTTPRequestsInflightMax()
	metrics.HttpRequestsInflightMax.WithLabelValues().Set(httpRequestsInflightMax)

	router := gin.New()
	router.Use(loggingMiddleware())
	router.Use(middleware.HTTPMetrics())
	router.Use(middleware.InflightRequests())
	router.GET("/code-2xx", handlers.Random2xxResponse)
	router.GET("/code-4xx", handlers.Random4xxResponse)
	router.GET("/code-5xx", handlers.Random5xxResponse)
	router.GET("/ms-200", handlers.LatencyHandler(200*time.Millisecond))
	router.GET("/ms-500", handlers.LatencyHandler(500*time.Millisecond))
	router.GET("/ms-1000", handlers.LatencyHandler(1000*time.Millisecond))

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	if err := router.Run(":8084"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func loggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(`{"time": "%s", "method": "%s", "path": "%s", "status": %d, "duration": %d}`+"\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency/time.Millisecond,
		)
	})
}
