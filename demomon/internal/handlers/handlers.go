package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Random2xxResponse(c *gin.Context) {
	statuses := []int{
		http.StatusOK,
		http.StatusAccepted,
		}
		index := rand.Intn(len(statuses) - 1)
		c.JSON(statuses[index], gin.H{"message": statuses[index]})
}

func Random4xxResponse(c *gin.Context) {
	statuses := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusNotFound,
		http.StatusTooManyRequests,
		http.StatusConflict,
		http.StatusTooManyRequests,
		}
		index := rand.Intn(len(statuses) - 1)
		c.JSON(statuses[index], gin.H{"message": fmt.Sprintf("%d response", statuses[index])})
}

func Random5xxResponse(c *gin.Context) {
		statuses := []int{
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
		http.StatusBadGateway,
	}

	index := rand.Intn(len(statuses) - 1)
	c.JSON(statuses[index], gin.H{"message": fmt.Sprintf("%d response", statuses[index])})
}

func LatencyHandler(latency time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		time.Sleep(latency)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d ms response", latency.Milliseconds()),
			"latency": time.Since(start).Milliseconds(),
		})
	}
}
