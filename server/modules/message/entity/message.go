package entity

import (
	"errors"
	"github.com/google/uuid"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/user_information/entity"
	"time"
)

type Message struct {
	Id             string     `json:"id,omitempty" gorm:"column:id;primaryKey"`
	ConversationId string     `json:"conversation_id,omitempty" gorm:"column:conversation_id"`
	SenderId       string     `json:"user_id,omitempty" gorm:"column:user_id"`
	Content        *string    `json:"message,omitempty" gorm:"column:message"`
	ReplyFor       *string    `json:"reply_for,omitempty" gorm:"column:reply_for"`
	SendingTime    *time.Time `json:"sending_time,omitempty" gorm:"column:sending_time"`
	//Message        *Message   `json:"reply_for,omitempty" gorm:"foreignKey:ReplyFor"`
	Sender entity.User `json:"user,omitempty" gorm:"foreignKey:SenderId;references:Id"`
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

var (
	ErrBlankContent   = app_error.ErrInvalidData(errors.New("content is blank"), "BLANK_CONTENT", "Content cannot be blank")
	ErrContentMissing = app_error.ErrInvalidData(errors.New("content is missing"), "CONTENT_MISSING", "content is required")
)
