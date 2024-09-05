package entity

import (
	"time"
)

type Conversation struct{
	Id string `json:"id" gorm:"column:id"`
	IsGroup bool `json:"is_group" gorm:"column:id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func NewConversation(id string, isGroup bool, createdAt time.Time) *Conversation{
	return &Conversation{
		Id: id,
		IsGroup: isGroup,
		CreatedAt: createdAt,
	}
}

func (c *Conversation)TableName() string{
	return "conversations"
}

type ConversationDetail struct{
	Id string `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Avatar string `json:"avatar" gorm:"column:avatar"`
	CreatedBy string `json:"created_by" gorm:"created_by"`
}

func (cd *ConversationDetail)TableName() string{
	return "conversation_details"
}

type ConversationRequest struct{
	SenderId string `json:"sender_id" gorm:"column:sender_id"`
	RecipientId string `json:"recipient_id" gorm:"column:recipient_id"`
	Status string `json:"status" gorm:"column:status"`
	RequestedTime *time.Time `json:"requested_time" gorm:"column:requested_time"`
	AcceptedTime *time.Time `json:"accepted_time" gorm:"column:accepted_time"`
}

func (cr *ConversationRequest)TableName() string{
	return "conversation_requests"
}

var (
	now = time.Now()
)

func NewConversationRequest(senderId, recipientId string) *ConversationRequest {
	return &ConversationRequest{
		SenderId:      senderId,
		RecipientId:   recipientId,
		RequestedTime: &now,
	}
}

func NewAcceptedConversationRequest() *ConversationRequest{
	return &ConversationRequest{
		Status:"ACCEPTED",
		AcceptedTime: &now,
	}
}
