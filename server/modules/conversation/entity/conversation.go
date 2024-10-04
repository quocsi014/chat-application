package entity

import (
	"errors"
	"github.com/quocsi014/common/app_error"
	"time"
)

type Conversation struct {
	Id              string     `json:"id,omitempty" gorm:"column:id"`
	IsGroup         bool       `json:"is_group,omitempty" gorm:"column:is_group"`
	LastMessageTime *time.Time `json:"last_message_time,omitempty" gorm:"column:last_message_time"`
	LastMessageId   *string    `json:"last_message_id,omitempty" gorm:"column:last_message_id"`
	CreatedAt       *time.Time `json:"created_at" gorm:"column:created_at"`
}

type ConversationResponse struct {
	Conversation
	ConversationDetail
	LastMessage    string `json:"message,omitempty" gorm:"column:message"`
	UserNameSender string `json:"user_name_sender,omitempty" gorm:"column:user_name_sender"`
}

func NewConversation(id string, isGroup bool) *Conversation {
	now := time.Now()
	return &Conversation{
		Id:        id,
		IsGroup:   isGroup,
		CreatedAt: &now,
	}
}

func (c *Conversation) TableName() string {
	return "conversations"
}

type ConversationDetail struct {
	Name      string `json:"name,omitempty" gorm:"column:name"`
	Avatar    string `json:"avatar_url,omitempty" gorm:"column:avatar_url"`
	CreatedBy string `json:"created_by,omitempty" gorm:"created_by"`
}

var (
	ErrBlankName   = app_error.ErrInvalidData(errors.New("Converstion name is blank"), "BLANK_NAME", "Conversation name cannot be blank")
	ErrNameMissing = app_error.ErrInvalidData(errors.New("Conversation name is missing"), "NAME_MISSING", "Conversation name is required")
)

func (cd *ConversationDetail) TableName() string {
	return "conversation_details"
}

type ConversationMembership struct {
	ConversationId string     `json:"conversation_id,omitempty" gorm:"column:conversation_id"`
	UserId         string     `json:"user_id,omitempty" gorm:"column:user_id"`
	Role           string     `json:"role,omitempty" gorm:"column:role"`
	JoinedTime     *time.Time `json:"joined_time,omitempty" gorm:"columnjoined_time"`
}

func (cm *ConversationMembership) TableName() string {
	return "conversation_memberships"
}

func NewConversationMembershipMemberRole(conversationId, userId string) *ConversationMembership {
	now := time.Now()
	return &ConversationMembership{
		ConversationId: conversationId,
		UserId:         userId,
		Role:           "MEMBER",
		JoinedTime:     &now,
	}
}
