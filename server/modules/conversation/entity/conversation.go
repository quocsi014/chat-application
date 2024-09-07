package entity

import (
	"time"
)

type Conversation struct{
	Id string `json:"id" gorm:"column:id"`
	IsGroup bool `json:"is_group" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func NewConversation(id string, isGroup bool) *Conversation{
	now := time.Now()
	return &Conversation{
		Id: id,
		IsGroup: isGroup,
		CreatedAt: &now,
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
	RequestedTime *time.Time `json:"requested_time" gorm:"column:requested_time"`
}

func (cr *ConversationRequest)TableName() string{
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

type ConversationMembership struct{
	ConversationId string `json:"conversation_id" gorm:"column:conversation_id"`
	UserId string `json:"user_id" gorm:"column:user_id"`
	Role string `json:"role" gorm:"column:role"`
	JoinedTime *time.Time `json:"joined_time" gorm:"columnjoined_time"`
}

func (cm *ConversationMembership)TableName() string {
	return "conversation_memberships"
}

func NewConversationMembershipMemberRole(conversationId, userId string) *ConversationMembership{
	now := time.Now()
	return &ConversationMembership{
		ConversationId: conversationId,
		UserId: userId,
		Role: "MEMBER",
		JoinedTime: &now,
	}
}
