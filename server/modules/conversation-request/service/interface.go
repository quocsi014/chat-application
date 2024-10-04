package service

import (
	"context"
	"github.com/quocsi014/modules/conversation-request/entity"
	conversationEntity "github.com/quocsi014/modules/conversation/entity"
)

type IConversationRequestRepository interface {
	CreateConversationRequest(ctx context.Context, req *entity.ConversationRequest) error
	AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*conversationEntity.Conversation, error)
	DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error
	GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error)
	GetConversationRequestReceived(ctx context.Context, recipientId string) ([]entity.ConversationRequestDetail, error)
}

type IConversationRequestService interface {
	CreateConversationRequest(ctx context.Context, senderId, recipientId string) error
	AcceptConversationRequest(ctx context.Context, senderId, recipientId string) (*conversationEntity.Conversation, error)
	DeleteConversationRequest(ctx context.Context, senderId, recipientId string) error
	GetConversationRequestSent(ctx context.Context, senderId string) ([]entity.ConversationRequestDetail, error)
	GetConversationRequestReceived(ctx context.Context, recipientId string) ([]entity.ConversationRequestDetail, error)
}
