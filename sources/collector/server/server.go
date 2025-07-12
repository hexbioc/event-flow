package server

import (
	"collector/config"
	"collector/events"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteGroupHandler interface {
	Attach(*gin.RouterGroup)
}

type routeGroupConfig struct {
	group   *gin.RouterGroup
	handler RouteGroupHandler
}

func New() *gin.Engine {
	// Read config
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Errorf("Fatal error while loading configuration: %w", err))
	}
	fmt.Println("config:", cfg)

	// Create server
	engine := gin.New()

	// Add common middleware
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	// Initialize middleware
	authMiddleware := AuthMiddleware(cfg)

	// Add healthcheck route
	engine.GET("/health", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	// Add events routes
	routeGroupHandlers := []routeGroupConfig{
		{engine.Group("/events", authMiddleware), &events.Handler{}},
	}

	for _, rgConfig := range routeGroupHandlers {
		rgConfig.handler.Attach(rgConfig.group)

	}

	return engine
}
