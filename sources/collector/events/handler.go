package events

import (
	"collector/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct{}

func (h *Handler) Attach(router *gin.RouterGroup) {
	router.POST("", h.createEvent)
}

func (h *Handler) createEvent(c *gin.Context) {
	logger := common.LoggerFromContext(c)
	var event CreateEventDTO

	if err := c.BindJSON(&event); err != nil {
		logger.Error("Unable to parse event body", zap.Error(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusAccepted)
}
