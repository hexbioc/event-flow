package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) Attach(router *gin.RouterGroup) {
	router.POST("", h.createEvent)
}

func (h *Handler) createEvent(c *gin.Context) {
	var event CreateEventDTO

	if err := c.BindJSON(&event); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusAccepted)
}
