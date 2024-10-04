package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quocsi014/common"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/helper"
	"github.com/quocsi014/middleware"
	"github.com/quocsi014/modules/message/entity"
	"github.com/quocsi014/modules/message/service"
	"net/http"
)

type MessageHandler struct {
	service service.IMessageService
}

func NewMessageHandler(service service.IMessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (mh *MessageHandler) SendMessage() func(*gin.Context) {
	return func(ctx *gin.Context) {
		message := entity.NewMessage()
		if err := ctx.ShouldBind(&message); err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}
		userId := helper.GetUserId(ctx)
		if userId == "" {
			return
		}
		message.SenderId = userId
		conversation_id := ctx.Param("conversation_id")
		message.ConversationId = conversation_id
		err := mh.service.SendMessage(ctx, message)
		if err != nil {
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, err)
			fmt.Println(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, message)
	}
}

func (mh *MessageHandler) GetMessages() func(*gin.Context) {
	return func(ctx *gin.Context) {
		paging := &(common.Paging{})
		if err := ctx.ShouldBind(paging); err != nil {
			ctx.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(err))
			return
		}
		paging.Process()
		conversationId := ctx.Param("conversation_id")

		messages, err := mh.service.GetMessages(ctx, paging, conversationId)
		if err != nil {
			fmt.Println(err.Error())
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"paging":   paging,
			"messages": messages,
		})
	}
}

func (mh *MessageHandler) SetupRoute(group *gin.RouterGroup) {
	group.POST("", middleware.VerifyToken(), mh.SendMessage())
	group.GET("", middleware.VerifyToken(), mh.GetMessages())
}
