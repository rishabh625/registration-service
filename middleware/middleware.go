package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LatencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		log.Printf("Method: %s, Path: %s, Status: %d, Latency: %v",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}
