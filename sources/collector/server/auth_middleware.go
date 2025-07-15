package server

import (
	"collector/config"
	"crypto/subtle"
	"net/http"

	"github.com/gin-gonic/gin"
)

const xApiKeyHeader = "X-Api-Key"

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	apiKeyBytes := []byte(cfg.XApiKey)

	return func(c *gin.Context) {
		headerValue := c.GetHeader(xApiKeyHeader)

		if subtle.ConstantTimeCompare([]byte(headerValue), apiKeyBytes) != 1 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
