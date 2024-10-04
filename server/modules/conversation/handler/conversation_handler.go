package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quocsi014/common"
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

func (h *ConversationHandler) GetConversations() func(*gin.Context) {
	return func(ctx *gin.Context) {

		userId := helper.GetUserId(ctx)
		if userId == "" {
			return
		}

		paging := common.PagingBinding(ctx)
		if paging == nil {
			return
		}

		conversations, err := h.service.GetConversations(userId, paging)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		pagingResponse := common.NewPagingResponse(paging, conversations)
		ctx.JSON(http.StatusOK, pagingResponse)
	}
}

func (h *ConversationHandler) SetupRoute(group *gin.RouterGroup) {
	group.GET("", middleware.VerifyToken(), h.GetConversations())
}
