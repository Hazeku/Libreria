package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware registra información sobre cada petición.
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Procesar la petición
		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		log.Printf("[GIN] %s %s | %d | %s | %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
			c.ClientIP(),
		)
	}
}