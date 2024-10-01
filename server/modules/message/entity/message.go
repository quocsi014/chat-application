package entity

import (
	"github.com/google/uuid"
	"github.com/quocsi014/modules/user_information/entity"
	"time"
)

type Message struct {
	Id             string     `json:"id" gorm:"column:id;primaryKey"`
	ConversationId string     `json:"conversation_id" gorm:"column:conversation_id"`
	SenderId       string     `json:"user_id" gorm:"column:user_id"`
	Content        *string    `json:"message" gorm:"column:message"`
	ReplyFor       *string    `json:"message" gorm:"column:reply_for"`
	SendingTime    *time.Time `json:"sending_time" gorm:"column:sending_time"`
	//Message        *Message   `json:"reply_for" gorm:"foreignKey:ReplyFor"`
	Sender entity.User `json:"user" gorm:"foreignKey:SenderId;references:Id"`
}

func NewMessage() *Message {
	message_id, _ := uuid.NewUUID()
	return &Message{
		Id: message_id.String(),
	}
}

func (m *Message) TableName() string {
	return "messages"
}
