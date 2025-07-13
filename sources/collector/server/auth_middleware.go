package server

import (
	"collector/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

const xApiKeyHeader = "X-Api-Key"

func AuthMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := c.GetHeader(xApiKeyHeader)

		if headerValue != config.XApiKey {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
