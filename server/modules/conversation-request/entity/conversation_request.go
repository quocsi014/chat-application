package entity

import (
	"github.com/quocsi014/modules/user_information/entity"
	"time"
)

type ConversationRequest struct {
	SenderId      string     `json:"sender_id" gorm:"column:sender_id"`
	RecipientId   string     `json:"recipient_id" gorm:"column:recipient_id"`
	RequestedTime *time.Time `json:"requested_time" gorm:"column:requested_time"`
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
	Sender    entity.User `json:"sender" gorm:"foreignKey:SenderId"`
	Recipient entity.User `json:"recipient" gorm:"foreignKey:RecipientId"`
}

func (crd *ConversationRequestDetail) TableName() string {
	return "conversation_requests"
}

type UserRelationship struct {
	UserId   string `json:"user_id" gorm:"column:user_id"`
	FriendId string `json:"friend_id" gorm:"column:friend_id"`
	Status   string `json:"status" gorm:"column:status"`
}

func NewUserRelationship(userId, friendId string) *UserRelationship {
	return &UserRelationship{
		UserId:   userId,
		FriendId: friendId,
		Status:   "PENDING",
	}
}

func NewUserRelationshipWithAccepted(userId, friendId string) *UserRelationship {
	return &UserRelationship{
		UserId:   userId,
		FriendId: friendId,
		Status:   "ACCEPTED",
	}
}

func (ur *UserRelationship) TableName() string {
	return "user_relationships"
}
