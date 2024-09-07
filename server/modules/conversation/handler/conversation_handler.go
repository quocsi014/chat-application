package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/helper"
	"github.com/quocsi014/middleware"
	"github.com/quocsi014/modules/conversation/entity"
)

type IConversationRequestService interface {
	CreateConversationRequest(ctx context.Context, senderId, recipientId string) error
	AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*entity.Conversation, error)
	DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error
	GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error)
	GetConversationRequestReceived(ctx context.Context, recipientId string) ([]entity.ConversationRequestDetail, error)
}

type ConversationRequestHandler struct {
	service IConversationRequestService
}

func NewConversationHandler(service IConversationRequestService) *ConversationRequestHandler {
	return &ConversationRequestHandler{
		service: service,
	}
}

func (crh *ConversationRequestHandler) CreateConversationRequest(c *gin.Context) {
	recipientId := c.Param("recipient_id")
	if recipientId == "" {
		c.JSON(http.StatusBadRequest, app_error.ErrInvalidRequest(nil))
		return
	}

	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Token không tồn tại"))
		return
	}
	
	claims, err := helper.GetMapClaims(token.(*jwt.Token))
	if err != nil {
		c.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(err, "Unauthorized"))
		return
	}

	senderId, ok := claims["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Invalid token"))
		return
	}

	if err := crh.service.CreateConversationRequest(c.Request.Context(), senderId, recipientId); err != nil {
		c.JSON(http.StatusInternalServerError, app_error.ErrInternal(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversation request sent successfully"})
}

func (crh *ConversationRequestHandler)AcceptConversationRequest() func(ctx *gin.Context){
	return func(ctx *gin.Context) {
		
		token, exists := ctx.Get("token")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Token không tồn tại"))
			return
		}
	
		claims, err := helper.GetMapClaims(token.(*jwt.Token))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(err, "Unauthorized"))
			return
		}

		recipientId, ok := claims["user_id"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, app_error.ErrUnauthenticatedError(nil, "Invalid token"))
			return
		}

		senderId := ctx.Param("sender_id")

		conversation, err := crh.service.AcceptConversationRequest(ctx, senderId, recipientId)
		if err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, err)
			return
		}
		ctx.Header("Location", "/" + conversation.Id)
		ctx.Status(http.StatusOK)
	}
}

func (crh *ConversationRequestHandler)RejectConversationRequest() func(*gin.Context){
	return func (ctx *gin.Context){
		recipientId := helper.GetUserId(ctx)	
		if recipientId == ""{
			return
		}
		senderId := ctx.Param("sender_id")
		err := crh.service.DeleteConversationRequest(ctx, senderId, recipientId)
		if err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, err)
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (crh *ConversationRequestHandler)DeleteConversationRequest() func(*gin.Context){
	return func (ctx *gin.Context){
		senderId := helper.GetUserId(ctx)	
		if senderId == ""{
			return
		}
		recipientId := ctx.Param("recipient_id")
		err := crh.service.DeleteConversationRequest(ctx, senderId, recipientId)
		if err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, err)
			return
		}
		ctx.Status(http.StatusOK)
	}
}

func (crh *ConversationRequestHandler) GetConversationRequestSent() func (*gin.Context){
	return func(ctx *gin.Context){
		senderId := helper.GetUserId(ctx)
		if senderId == ""{
			return
		}
		conversationReqs, err := crh.service.GetConversationRequestSent(ctx, senderId)
		if err != nil{
			errResponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errResponse.Code, err)
			return
		}
		ctx.JSON(http.StatusOK, conversationReqs)
	}
}

func (crh *ConversationRequestHandler) GetConversationRequestReceived() func(*gin.Context){
	return func(ctx *gin.Context) {
		recipientId := helper.GetUserId(ctx)
		if recipientId == ""{
			return
		}
		conversationReqs, err := crh.service.GetConversationRequestReceived(ctx, recipientId)
		if err != nil{
			errReponse := app_error.NewErrorResponseWithAppError(err)
			ctx.JSON(errReponse.Code, err)
			return
		}
		ctx.JSON(http.StatusOK, conversationReqs)
	}
}
func (crh *ConversationRequestHandler) SetupRoute(group *gin.RouterGroup) {
	group.POST("/requests/sent/:recipient_id", middleware.VerifyToken(), crh.CreateConversationRequest)
	group.POST("/requests/received/:sender_id/accept", middleware.VerifyToken(), crh.AcceptConversationRequest())
	group.POST("/requests/received/:sender_id/reject", middleware.VerifyToken(), crh.RejectConversationRequest())
	group.DELETE("/requests/sent/:recipient_id", middleware.VerifyToken(), crh.DeleteConversationRequest())
	group.GET("requests/sent", middleware.VerifyToken(), crh.GetConversationRequestSent())
	group.GET("requests/received", middleware.VerifyToken(), crh.GetConversationRequestReceived())
}


