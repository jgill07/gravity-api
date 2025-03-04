package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jgill07/gravity-api/internal/log"
	"go.uber.org/zap"
)

func logger(c *gin.Context) {
	t := time.Now()
	c.Next()
	latency := time.Since(t)

	status := c.Writer.Status()
	method := c.Request.Method
	path := c.Request.URL.Path

	log.WithFields(
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status", status),
		zap.Duration("latency", latency),
	).Debug("handled request")
}
