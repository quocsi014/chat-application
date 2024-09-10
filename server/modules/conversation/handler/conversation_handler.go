package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quocsi014/helper"
	"github.com/quocsi014/middleware"
	"github.com/quocsi014/modules/conversation/service"
)

type ConversationHandler struct {
	service service.IConversationService
}

func NewConversationHandler(service service.IConversationService) *ConversationHandler {
	return &ConversationHandler{service: service}
}

func (h *ConversationHandler) GetConversations(c *gin.Context) {
	userId := helper.GetUserId(c)
	if userId == "" {
		return
	}

	conversations, err := h.service.GetConversations(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, conversations)
}

func (h *ConversationHandler) SetupRoute(group *gin.RouterGroup) {
	group.GET("", middleware.VerifyToken(), h.GetConversations)
}
