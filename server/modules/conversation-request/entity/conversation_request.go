package entity

import (
	"github.com/quocsi014/modules/user_information/entity"
	"time"
)

type ConversationRequest struct {
	SenderId      string     `json:"sender_id,omitempty" gorm:"column:sender_id"`
	RecipientId   string     `json:"recipient_id,omitempty" gorm:"column:recipient_id"`
	RequestedTime *time.Time `json:"requested_time,omitempty" gorm:"column:requested_time"`
}

func (cr *ConversationRequest) TableName() string {
	return "conversation_requests"
}

func NewConversationRequest(senderId, recipientId string) *ConversationRequest {
	now := time.Now()
	return &ConversationRequest{
		SenderId:      senderId,
		RecipientId:   recipientId,
		RequestedTime: &now,
	}
}

type ConversationRequestDetail struct {
	ConversationRequest
	Sender    entity.User `json:"sender,omitempty" gorm:"foreignKey:SenderId"`
	Recipient entity.User `json:"recipient,omitempty" gorm:"foreignKey:RecipientId"`
}

func (crd *ConversationRequestDetail) TableName() string {
	return "conversation_requests"
}
