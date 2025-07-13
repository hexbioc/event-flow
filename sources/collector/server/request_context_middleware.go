package server

import (
	"collector/common"
	"collector/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const xRequestIdHeader = "X-Request-Id"

func RequestContextMiddleware(config *config.Config, baseLogger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader(xRequestIdHeader)
		if requestId == "" {
			requestId = uuid.NewString()
		}

		// Create logger with reqeust-scoped fields
		logger := baseLogger.With(zap.String("request_id", requestId))

		// Update context with request-scoped fields
		c.Set("xRequestId", requestId)

		// Attach logger to context for use downstream
		common.AddLoggerToContext(c, logger)
	}
}
