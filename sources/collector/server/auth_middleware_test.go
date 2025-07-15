package server

import (
	"collector/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	r := createDummyGinServer()

	cfg := config.Config{}
	cfg.XApiKey = "dummyKey"

	// Add a protected route
	r.GET("/protected", AuthMiddleware(&cfg), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Test with incorrect key
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("X-Api-Key", "incorrectKey")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test with valid key
	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("X-Api-Key", cfg.XApiKey)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
