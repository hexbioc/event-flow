package server

import (
	"collector/common"
	"collector/config"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerMiddleware(config *config.Config, skipRoutes []string) gin.HandlerFunc {
	skipRoutesMap := make(map[string]bool, len(skipRoutes))
	for _, skipRoute := range skipRoutes {
		skipRoutesMap[skipRoute] = true
	}

	return func(c *gin.Context) {
		start := time.Now()

		// Check if route logging needs to be skipped
		requestPath := c.Request.URL.Path

		if _, ok := skipRoutesMap[requestPath]; ok {
			return
		}

		// Proceed with the request
		c.Next()

		// Setup fields for logging and log request
		logger := common.LoggerFromContext(c)

		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()
		fields := []zapcore.Field{
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", requestPath),
			zap.String("ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if status >= 400 {
			logger.Error(requestPath, fields...)
		} else {
			logger.Info(requestPath, fields...)
		}
	}
}
