package common

import (
	"collector/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const loggerContextKey = "requestScopedLogger"

func NewLogger(cfg *config.Config) *zap.Logger {
	var logger *zap.Logger
	var loggerConfig zap.Config

	if cfg.Env == "dev" {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
	}

	loggerConfig.DisableStacktrace = true

	logger, err := loggerConfig.Build()
	if err != nil {
		panic("Failed to create logger")
	}

	return logger
}

func AddLoggerToContext(c *gin.Context, logger *zap.Logger) {
	c.Set(loggerContextKey, logger)
}

func LoggerFromContext(c *gin.Context) *zap.Logger {
	// We assume that the logger will always be a part of context
	logger, _ := c.Value(loggerContextKey).(*zap.Logger)
	return logger
}
