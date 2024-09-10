package entity

import (
	"time"

	"github.com/quocsi014/modules/user_information/entity"
)

type MessageInformation struct {
	ConversationId string     `json:"conversation_id" gorm:"column:conversation_id"`
	UserId         string     `json:"user_id" gorm:"column:user_id"`
	Message        *string    `json:"message" gorm:"column:message"`
	ReplyFor       *string    `json:"reply_for" gorm:"column:reply_for"`
	SendingTime    *time.Time `json:"sending_time" gorm:"column:sending_time"`
}

type Message struct {
	Id string `json:"id" gorm:"column:id;primaryKey"`
	MessageInformation
}

func (m *Message) TableName() string {
	return "message"
}


type MessageResponse struct {
	Message
	entity.UserInformation
}