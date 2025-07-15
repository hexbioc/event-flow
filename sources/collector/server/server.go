package server

import (
	"collector/common"
	"collector/config"
	"collector/events"
	"collector/messaging"
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

func New() (engine *gin.Engine, shutdownActions func()) {
	// Read config
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Errorf("Fatal error while loading configuration: %w", err))
	}

	// Setup RabbitMQ
	rmqHandler := messaging.New(cfg)
	rmqConnection, rmqChannel := rmqHandler.Connect()

	// Create server
	if cfg.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine = gin.New()

	// Create logger
	baseLogger := common.NewLogger(cfg)

	// Add common middleware
	engine.Use(RequestContextMiddleware(cfg, baseLogger))
	engine.Use(LoggerMiddleware(cfg, []string{"/health"}))
	engine.Use(RecoveryMiddleware())

	// Initialize route-based middleware
	authMiddleware := AuthMiddleware(cfg)

	// Add healthcheck route
	engine.GET("/health", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	// Add events routes
	routeGroupHandlers := []routeGroupConfig{
		{engine.Group("/events", authMiddleware), &events.Handler{Mq: rmqHandler}},
	}

	for _, rgConfig := range routeGroupHandlers {
		rgConfig.handler.Attach(rgConfig.group)

	}

	shutdownActions = func() {
		defer rmqConnection.Close()
		defer rmqChannel.Close()
	}

	return engine, shutdownActions
}
