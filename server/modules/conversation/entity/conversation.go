package entity

import (
	"time"
)

type Conversation struct{
	Id string `json:"id" gorm:"column:id"`
	IsGroup bool `json:"is_group" gorm:"column:is_group"`
	LastMessageTime *time.Time `json:"last_message_time" gorm:"column:last_message_time"`
	LastMessageId *string `json:"last_message_id" gorm:"column:last_message_id"`
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
	Name string `json:"name" gorm:"column:name"`
	Avatar string `json:"avatar_url" gorm:"column:avatar_url"`
	CreatedBy string `json:"created_by" gorm:"created_by"`
}

func (cd *ConversationDetail)TableName() string{
	return "conversation_details"
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

type ConversationResponse struct{
	Conversation
	ConversationDetail
	LastMessage string`json:"message" gorm:"column:message"`
	UserNameSender string `json:"user_name_sender" gorm:"column:user_name_sender"`
}
