package events

import (
	"collector/common"
	"collector/messaging"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	Mq *messaging.Handler
}

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

	// Publish message to queue, log on failure
	err := h.Mq.PublishJSON(event)
	if err != nil {
		eventJsonBytes, _ := json.Marshal(event)
		logger.Error(
			"Failed to publish event to queue",
			zap.String("event", string(eventJsonBytes)),
			zap.Error(err),
		)
	}

	c.Status(http.StatusAccepted)
}
